package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/blogs"
	common "earnforglance/server/domain/common"

	"github.com/gin-gonic/gin"
)

type BlogSettingsController struct {
	BlogSettingsUsecase domain.BlogSettingsUsecase
	Env                 *bootstrap.Env
}

func (tc *BlogSettingsController) Create(c *gin.Context) {
	var task domain.BlogSettings
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

	err = tc.BlogSettingsUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "BlogSettings created successfully",
	})
}

func (tc *BlogSettingsController) Update(c *gin.Context) {
	var task domain.BlogSettings
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

	err = tc.BlogSettingsUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "BlogSettings update successfully",
	})
}

func (tc *BlogSettingsController) Delete(c *gin.Context) {
	var task domain.BlogSettings
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

	err = tc.BlogSettingsUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "BlogSettings update successfully",
	})
}

func (lc *BlogSettingsController) FetchByID(c *gin.Context) {
	BlogSettingsID := c.Query("id")
	if BlogSettingsID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	BlogSettings, err := lc.BlogSettingsUsecase.FetchByID(c, BlogSettingsID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: BlogSettingsID})
		return
	}

	c.JSON(http.StatusOK, BlogSettings)
}

func (lc *BlogSettingsController) Fetch(c *gin.Context) {

	BlogSettings, err := lc.BlogSettingsUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, BlogSettings)
}
