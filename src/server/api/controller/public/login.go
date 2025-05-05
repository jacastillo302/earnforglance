package controller

import (
	"context"
	"crypto/rand"
	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	custumers "earnforglance/server/domain/customers"
	domain "earnforglance/server/domain/public"
	security "earnforglance/server/domain/security"
	service "earnforglance/server/service/security"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "https://personasyrecursos.com/login_google_redirect",
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

// GoogleClaims represents the claims returned by Google OAuth.
type GoogleClaims struct {
	Aud           string `json:"aud"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Exp           int64  `json:"exp"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Iat           int64  `json:"iat"`
	Iss           string `json:"iss"`
	Locale        string `json:"locale"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Sub           string `json:"sub"`
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func (lc *LoginController) LoginGoogle(c *gin.Context) {

	// Create oauthState cookie
	oauthState := generateStateOauthCookie(c.Writer)

	googleOauthConfig.ClientID = lc.Env.GoogleOauthClientID
	googleOauthConfig.ClientSecret = lc.Env.GoogleOauthClientSecret

	u := googleOauthConfig.AuthCodeURL(oauthState)

	fmt.Println("oauthState: ", u)

	http.Redirect(c.Writer, c.Request, u, http.StatusTemporaryRedirect)

}

func (lc *LoginController) LoginGoogleRedirec(c *gin.Context) {
	// Read oauthState from Cookie

	oauthState, _ := c.Request.Cookie("oauthstate")

	if lc.Env.AppEnv == "development" {
		if oauthState == nil {
			oauthState = &http.Cookie{
				Name:  "oauthstate",
				Value: c.Query("state"),
			}
		}
	}

	if c.Query("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getUserDataFromGoogle(c.Query("code"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: err.Error()})
		return
	}

	var response GoogleClaims

	err = json.Unmarshal(data, &response)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: err.Error()})
		return
	}

	// GetOrCreate User in your db.
	// Redirect or response with a token.
	// More code .....
	fmt.Fprintf(c.Writer, "UserInfo: %s\n", data)

	c.JSON(http.StatusOK, response)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	// Use code to get token and get user info from Google.

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}

func (lc *LoginController) Login(c *gin.Context) {
	var request domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: err.Error()})
		return
	}

	request.Lang = c.Query("language")
	request.Email = c.Query("email")
	request.Password = c.Query("password")

	sLang := ""
	if request.Lang == "" {
		lang, err := lc.LoginUsecase.GetSettingByName(c, "DefaultAdminLanguageId")
		if err != nil {
			c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
			return
		}

		sLang = lang.Value
	} else {
		lang, err := lc.LoginUsecase.GetLangugaByCode(c, request.Lang)
		if err != nil {
			c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
			return
		}
		sLang = lang.ID.Hex()
	}

	setting, err := lc.LoginUsecase.GetSettingByName(c, "UsernamesEnabled")
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	var login custumers.Customer
	if setting.Value == "true" {
		user, err := lc.LoginUsecase.GetByUserName(c, request.Email)
		if err != nil {
			locale, err := lc.LoginUsecase.GetLocalebyName(c, "Account.Login.WrongCredentials.CustomerNotExist", sLang)
			if err != nil {
				c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
				return
			}
			c.JSON(http.StatusNotFound, common.ErrorResponse{Message: locale.ResourceValue})
			return
		}

		login = user

	} else {

		user, err := lc.LoginUsecase.GetUserByEmail(c, request.Email)
		if err != nil {
			locale, err := lc.LoginUsecase.GetLocalebyName(c, "Account.PasswordRecovery.EmailNotFound", sLang)
			if err != nil {
				c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
				return
			}
			c.JSON(http.StatusNotFound, common.ErrorResponse{Message: locale.ResourceValue})
			return
		}

		login = user
	}

	encripkey, err := lc.LoginUsecase.GetSettingByName(c, "EncryptionKey")
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	algot, err := lc.LoginUsecase.GetSettingByName(c, "UseAesEncryptionAlgorithm")
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	algotBool, err := strconv.ParseBool(algot.Value)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Invalid value for UseAesEncryptionAlgorithm setting"})
		return
	}

	pformat, err := lc.LoginUsecase.GetSettingByName(c, "HashedPasswordFormat")
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	valresult := service.ValidateCustomer(login)
	if valresult != "Successful" {
		locale, err := lc.LoginUsecase.GetLocalebyName(c, valresult, sLang)
		if err != nil {
			c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
			return
		}
		c.JSON(http.StatusUnauthorized, common.ErrorResponse{Message: locale.ResourceValue})
		return
	}

	multif, err := lc.LoginUsecase.GetSettingByName(c, "ForceMultifactorAuthentication")
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	multifBool, err := strconv.ParseBool(multif.Value)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Invalid value for ForceMultifactorAuthentication setting"})
		return
	}

	if multifBool {
		locale, err := lc.LoginUsecase.GetLocalebyName(c, "Account.MultiFactorAuthentication.Warning.ForceActivation", sLang)
		if err != nil {
			c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
			return
		}
		c.JSON(http.StatusUnauthorized, common.ErrorResponse{Message: locale.ResourceValue})
		return
	}

	psw, err := lc.LoginUsecase.GetPasw(c, login.ID.Hex())
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	if !service.PasswordsMatch(psw.PasswordFormatID, psw.Password, psw.PasswordSalt, request.Password, encripkey.Value, algotBool, pformat.Value) {

		locale, err := lc.LoginUsecase.GetLocalebyName(c, "Account.Login.WrongCredentials", sLang)
		if err != nil {
			c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
			return
		}

		// Insert activity log WrongCredentials
		_, err = lc.LoginUsecase.AddActivityLog(c, login.ID, "PublicStore.Login", locale.ResourceValue, login.LastIpAddress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusUnauthorized, common.ErrorResponse{Message: locale.ResourceValue})
		return
	}

	locale, err := lc.LoginUsecase.GetLocalebyName(c, "ActivityLog.PublicStore.Login", sLang)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	// Insert activity log Login
	_, err = lc.LoginUsecase.AddActivityLog(c, login.ID, "PublicStore.Login", locale.ResourceValue, login.LastIpAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	slugs := []security.UrlRecord{{
		Name:     "Customer",
		IsRead:   true,
		IsWrite:  true,
		IsDelete: true,
		IsUpdate: true,
	}}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&login, slugs, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&login, slugs, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
