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

type ShippingMethodController struct {
	ShippingMethodUsecase domain.ShippingMethodUsecase
	Env                   *bootstrap.Env
}

func (tc *ShippingMethodController) Create(c *gin.Context) {
	var task domain.ShippingMethod
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

	err = tc.ShippingMethodUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ShippingMethod created successfully",
	})
}

func (tc *ShippingMethodController) Update(c *gin.Context) {
	var task domain.ShippingMethod
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

	err = tc.ShippingMethodUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ShippingMethod update successfully",
	})
}

func (tc *ShippingMethodController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ShippingMethodUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ShippingMethodController) FetchByID(c *gin.Context) {
	ShippingMethodID := c.Query("id")
	if ShippingMethodID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ShippingMethod, err := lc.ShippingMethodUsecase.FetchByID(c, ShippingMethodID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ShippingMethodID})
		return
	}

	c.JSON(http.StatusOK, ShippingMethod)
}

func (lc *ShippingMethodController) Fetch(c *gin.Context) {

	ShippingMethod, err := lc.ShippingMethodUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ShippingMethod)
}
