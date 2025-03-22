package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/catalog"
	common "earnforglance/server/domain/common"

	"github.com/gin-gonic/gin"
)

type ProductEditorSettingsController struct {
	ProductEditorSettingsUsecase domain.ProductEditorSettingsUsecase
	Env                          *bootstrap.Env
}

func (tc *ProductEditorSettingsController) Create(c *gin.Context) {
	var task domain.ProductEditorSettings
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

	err = tc.ProductEditorSettingsUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductEditorSettings created successfully",
	})
}

func (tc *ProductEditorSettingsController) Update(c *gin.Context) {
	var task domain.ProductEditorSettings
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

	err = tc.ProductEditorSettingsUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductEditorSettings update successfully",
	})
}

func (tc *ProductEditorSettingsController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ProductEditorSettingsUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ProductEditorSettingsController) FetchByID(c *gin.Context) {
	ProductEditorSettingsID := c.Query("id")
	if ProductEditorSettingsID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ProductEditorSettings, err := lc.ProductEditorSettingsUsecase.FetchByID(c, ProductEditorSettingsID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ProductEditorSettingsID})
		return
	}

	c.JSON(http.StatusOK, ProductEditorSettings)
}

func (lc *ProductEditorSettingsController) Fetch(c *gin.Context) {

	ProductEditorSettings, err := lc.ProductEditorSettingsUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProductEditorSettings)
}
