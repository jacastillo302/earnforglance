package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/localization"

	"github.com/gin-gonic/gin"
)

type LocaleStringResourceController struct {
	LocaleStringResourceUsecase domain.LocaleStringResourceUsecase
	Env                         *bootstrap.Env
}

func (tc *LocaleStringResourceController) Create(c *gin.Context) {
	var task domain.LocaleStringResource
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

	err = tc.LocaleStringResourceUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "LocaleStringResource created successfully",
	})
}

func (tc *LocaleStringResourceController) Update(c *gin.Context) {
	var task domain.LocaleStringResource
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

	err = tc.LocaleStringResourceUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "LocaleStringResource update successfully",
	})
}

func (tc *LocaleStringResourceController) Delete(c *gin.Context) {
	var task domain.LocaleStringResource
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

	err = tc.LocaleStringResourceUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "LocaleStringResource update successfully",
	})
}

func (lc *LocaleStringResourceController) FetchByID(c *gin.Context) {
	LocaleStringResourceID := c.Query("id")
	if LocaleStringResourceID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	LocaleStringResource, err := lc.LocaleStringResourceUsecase.FetchByID(c, LocaleStringResourceID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: LocaleStringResourceID})
		return
	}

	c.JSON(http.StatusOK, LocaleStringResource)
}

func (lc *LocaleStringResourceController) Fetch(c *gin.Context) {

	LocaleStringResource, err := lc.LocaleStringResourceUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, LocaleStringResource)
}
