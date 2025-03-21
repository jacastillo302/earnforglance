package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/gdpr"

	"github.com/gin-gonic/gin"
)

type GdprConsentController struct {
	GdprConsentUsecase domain.GdprConsentUsecase
	Env                *bootstrap.Env
}

func (tc *GdprConsentController) Create(c *gin.Context) {
	var task domain.GdprConsent
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

	err = tc.GdprConsentUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "GdprConsent created successfully",
	})
}

func (tc *GdprConsentController) Update(c *gin.Context) {
	var task domain.GdprConsent
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

	err = tc.GdprConsentUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "GdprConsent update successfully",
	})
}

func (tc *GdprConsentController) Delete(c *gin.Context) {
	var task domain.GdprConsent
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

	err = tc.GdprConsentUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "GdprConsent update successfully",
	})
}

func (lc *GdprConsentController) FetchByID(c *gin.Context) {
	GdprConsentID := c.Query("id")
	if GdprConsentID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	GdprConsent, err := lc.GdprConsentUsecase.FetchByID(c, GdprConsentID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: GdprConsentID})
		return
	}

	c.JSON(http.StatusOK, GdprConsent)
}

func (lc *GdprConsentController) Fetch(c *gin.Context) {

	GdprConsent, err := lc.GdprConsentUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GdprConsent)
}
