package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"

	"github.com/gin-gonic/gin"
)

type AddressAttributeValueController struct {
	AddressAttributeValueUsecase common.AddressAttributeValueUsecase
	Env                          *bootstrap.Env
}

func (tc *AddressAttributeValueController) Create(c *gin.Context) {
	var task common.AddressAttributeValue
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

	err = tc.AddressAttributeValueUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "AddressAttributeValue created successfully",
	})
}

func (tc *AddressAttributeValueController) Update(c *gin.Context) {
	var task common.AddressAttributeValue
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

	err = tc.AddressAttributeValueUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "AddressAttributeValue update successfully",
	})
}

func (tc *AddressAttributeValueController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.AddressAttributeValueUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *AddressAttributeValueController) FetchByID(c *gin.Context) {
	AddressAttributeValueID := c.Query("id")
	if AddressAttributeValueID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	AddressAttributeValue, err := lc.AddressAttributeValueUsecase.FetchByID(c, AddressAttributeValueID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: AddressAttributeValueID})
		return
	}

	c.JSON(http.StatusOK, AddressAttributeValue)
}

func (lc *AddressAttributeValueController) Fetch(c *gin.Context) {

	AddressAttributeValue, err := lc.AddressAttributeValueUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, AddressAttributeValue)
}
