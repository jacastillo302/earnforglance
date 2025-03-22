package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/shipping"

	"github.com/gin-gonic/gin"
)

type ShippingOptionController struct {
	ShippingOptionUsecase domain.ShippingOptionUsecase
	Env                   *bootstrap.Env
}

func (tc *ShippingOptionController) Create(c *gin.Context) {
	var task domain.ShippingOption
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

	err = tc.ShippingOptionUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ShippingOption created successfully",
	})
}

func (tc *ShippingOptionController) Update(c *gin.Context) {
	var task domain.ShippingOption
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

	err = tc.ShippingOptionUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ShippingOption update successfully",
	})
}

func (tc *ShippingOptionController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ShippingOptionUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ShippingOptionController) FetchByID(c *gin.Context) {
	ShippingOptionID := c.Query("id")
	if ShippingOptionID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ShippingOption, err := lc.ShippingOptionUsecase.FetchByID(c, ShippingOptionID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ShippingOptionID})
		return
	}

	c.JSON(http.StatusOK, ShippingOption)
}

func (lc *ShippingOptionController) Fetch(c *gin.Context) {

	ShippingOption, err := lc.ShippingOptionUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ShippingOption)
}
