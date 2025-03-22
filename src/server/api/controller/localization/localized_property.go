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

type LocalizedPropertyController struct {
	LocalizedPropertyUsecase domain.LocalizedPropertyUsecase
	Env                      *bootstrap.Env
}

func (tc *LocalizedPropertyController) Create(c *gin.Context) {
	var task domain.LocalizedProperty
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

	err = tc.LocalizedPropertyUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "LocalizedProperty created successfully",
	})
}

func (tc *LocalizedPropertyController) Update(c *gin.Context) {
	var task domain.LocalizedProperty
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

	err = tc.LocalizedPropertyUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "LocalizedProperty update successfully",
	})
}

func (tc *LocalizedPropertyController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.LocalizedPropertyUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *LocalizedPropertyController) FetchByID(c *gin.Context) {
	LocalizedPropertyID := c.Query("id")
	if LocalizedPropertyID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	LocalizedProperty, err := lc.LocalizedPropertyUsecase.FetchByID(c, LocalizedPropertyID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: LocalizedPropertyID})
		return
	}

	c.JSON(http.StatusOK, LocalizedProperty)
}

func (lc *LocalizedPropertyController) Fetch(c *gin.Context) {

	LocalizedProperty, err := lc.LocalizedPropertyUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, LocalizedProperty)
}
