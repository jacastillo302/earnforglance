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

type CheckoutAttributeValueController struct {
	CheckoutAttributeValueUsecase domain.CheckoutAttributeValueUsecase
	Env                           *bootstrap.Env
}

func (tc *CheckoutAttributeValueController) Create(c *gin.Context) {
	var task domain.CheckoutAttributeValue
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

	err = tc.CheckoutAttributeValueUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CheckoutAttributeValue created successfully",
	})
}

func (tc *CheckoutAttributeValueController) Update(c *gin.Context) {
	var task domain.CheckoutAttributeValue
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

	err = tc.CheckoutAttributeValueUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CheckoutAttributeValue update successfully",
	})
}

func (tc *CheckoutAttributeValueController) Delete(c *gin.Context) {
	var task domain.CheckoutAttributeValue
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

	err = tc.CheckoutAttributeValueUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CheckoutAttributeValue update successfully",
	})
}

func (lc *CheckoutAttributeValueController) FetchByID(c *gin.Context) {
	CheckoutAttributeValueID := c.Query("id")
	if CheckoutAttributeValueID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	CheckoutAttributeValue, err := lc.CheckoutAttributeValueUsecase.FetchByID(c, CheckoutAttributeValueID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: CheckoutAttributeValueID})
		return
	}

	c.JSON(http.StatusOK, CheckoutAttributeValue)
}

func (lc *CheckoutAttributeValueController) Fetch(c *gin.Context) {

	CheckoutAttributeValue, err := lc.CheckoutAttributeValueUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CheckoutAttributeValue)
}
