package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/messages"

	"github.com/gin-gonic/gin"
)

type EmailAccountController struct {
	EmailAccountUsecase domain.EmailAccountUsecase
	Env                 *bootstrap.Env
}

func (tc *EmailAccountController) Create(c *gin.Context) {
	var task domain.EmailAccount
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Failed to read request body"})
		return
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Invalid request body"})
		return
	}

	err = tc.EmailAccountUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "EmailAccount created successfully",
	})
}

func (tc *EmailAccountController) Update(c *gin.Context) {
	var task domain.EmailAccount
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Failed to read request body"})
		return
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Invalid request body"})
		return
	}

	err = tc.EmailAccountUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "EmailAccount update successfully",
	})
}

func (tc *EmailAccountController) Delete(c *gin.Context) {
	var task domain.EmailAccount
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Failed to read request body"})
		return
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Invalid request body"})
		return
	}

	err = tc.EmailAccountUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "EmailAccount update successfully",
	})
}

func (lc *EmailAccountController) FetchByID(c *gin.Context) {
	EmailAccountID := c.Query("id")
	if EmailAccountID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	EmailAccount, err := lc.EmailAccountUsecase.FetchByID(c, EmailAccountID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: EmailAccountID})
		return
	}

	c.JSON(http.StatusOK, EmailAccount)
}

func (lc *EmailAccountController) Fetch(c *gin.Context) {

	EmailAccount, err := lc.EmailAccountUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, EmailAccount)
}
