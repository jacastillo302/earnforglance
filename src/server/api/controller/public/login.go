package controller

import (
	"net/http"
	"strconv"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	custumers "earnforglance/server/domain/customers"
	domain "earnforglance/server/domain/public"
	security "earnforglance/server/domain/security"
	service "earnforglance/server/service/security"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
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
