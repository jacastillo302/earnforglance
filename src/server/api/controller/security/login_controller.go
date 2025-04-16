package controller

import (
	"net/http"
	"strconv"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	custumers "earnforglance/server/domain/customers"
	domain "earnforglance/server/domain/security"
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

	setting, err := lc.LoginUsecase.GetSettingByName(c, "UsernamesEnabled")
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	locale, err := lc.LoginUsecase.GetLocalebyName(c, "Account.PasswordRecovery.EmailNotFound", "64f1a2c3d4e5f67890123456")
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	var login custumers.Customer
	if setting.Value == "Enabled" {
		user, err := lc.LoginUsecase.GetByUserName(c, request.Email)
		if err != nil {
			c.JSON(http.StatusNotFound, common.ErrorResponse{Message: locale.ResourceValue})
			return
		}

		login = user

	} else {
		user, err := lc.LoginUsecase.GetUserByEmail(c, request.Email)
		if err != nil {
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
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: valresult})
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
		c.JSON(http.StatusUnauthorized, common.ErrorResponse{Message: "Multifactor authentication required"})
		return
	}

	psw, err := lc.LoginUsecase.GetPasw(c, login.ID.Hex())
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	if !service.PasswordsMatch(psw.PasswordFormatID, psw.Password, psw.PasswordSalt, request.Password, encripkey.Value, algotBool, pformat.Value) {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: "The Password is worng"})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&login, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&login, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
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
