package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"

	"github.com/gin-gonic/gin"
)

type AddressAttributeController struct {
	AddressAttributeUsecase common.AddressAttributeUsecase
	Env                     *bootstrap.Env
}

func (tc *AddressAttributeController) Create(c *gin.Context) {
	var task common.AddressAttribute
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

	err = tc.AddressAttributeUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "AddressAttribute created successfully",
	})
}

func (tc *AddressAttributeController) Update(c *gin.Context) {
	var task common.AddressAttribute
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

	err = tc.AddressAttributeUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "AddressAttribute update successfully",
	})
}

func (tc *AddressAttributeController) Delete(c *gin.Context) {
	var task common.AddressAttribute
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

	err = tc.AddressAttributeUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "AddressAttribute update successfully",
	})
}

func (lc *AddressAttributeController) FetchByID(c *gin.Context) {
	AddressAttributeID := c.Query("id")
	if AddressAttributeID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	AddressAttribute, err := lc.AddressAttributeUsecase.FetchByID(c, AddressAttributeID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: AddressAttributeID})
		return
	}

	c.JSON(http.StatusOK, AddressAttribute)
}

func (lc *AddressAttributeController) Fetch(c *gin.Context) {

	AddressAttribute, err := lc.AddressAttributeUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, AddressAttribute)
}
