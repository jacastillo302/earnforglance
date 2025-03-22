package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/orders"

	"github.com/gin-gonic/gin"
)

type ShoppingCartSettingsController struct {
	ShoppingCartSettingsUsecase domain.ShoppingCartSettingsUsecase
	Env                         *bootstrap.Env
}

func (tc *ShoppingCartSettingsController) Create(c *gin.Context) {
	var task domain.ShoppingCartSettings
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

	err = tc.ShoppingCartSettingsUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ShoppingCartSettings created successfully",
	})
}

func (tc *ShoppingCartSettingsController) Update(c *gin.Context) {
	var task domain.ShoppingCartSettings
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

	err = tc.ShoppingCartSettingsUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ShoppingCartSettings update successfully",
	})
}

func (tc *ShoppingCartSettingsController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ShoppingCartSettingsUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ShoppingCartSettingsController) FetchByID(c *gin.Context) {
	ShoppingCartSettingsID := c.Query("id")
	if ShoppingCartSettingsID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ShoppingCartSettings, err := lc.ShoppingCartSettingsUsecase.FetchByID(c, ShoppingCartSettingsID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ShoppingCartSettingsID})
		return
	}

	c.JSON(http.StatusOK, ShoppingCartSettings)
}

func (lc *ShoppingCartSettingsController) Fetch(c *gin.Context) {

	ShoppingCartSettings, err := lc.ShoppingCartSettingsUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ShoppingCartSettings)
}
