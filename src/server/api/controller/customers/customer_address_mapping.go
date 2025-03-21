package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/customers"

	"github.com/gin-gonic/gin"
)

type CustomerAddressMappingController struct {
	CustomerAddressMappingUsecase domain.CustomerAddressMappingUsecase
	Env                           *bootstrap.Env
}

func (tc *CustomerAddressMappingController) Create(c *gin.Context) {
	var task domain.CustomerAddressMapping
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

	err = tc.CustomerAddressMappingUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CustomerAddressMapping created successfully",
	})
}

func (tc *CustomerAddressMappingController) Update(c *gin.Context) {
	var task domain.CustomerAddressMapping
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

	err = tc.CustomerAddressMappingUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CustomerAddressMapping update successfully",
	})
}

func (tc *CustomerAddressMappingController) Delete(c *gin.Context) {
	var task domain.CustomerAddressMapping
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

	err = tc.CustomerAddressMappingUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CustomerAddressMapping update successfully",
	})
}

func (lc *CustomerAddressMappingController) FetchByID(c *gin.Context) {
	CustomerAddressMappingID := c.Query("id")
	if CustomerAddressMappingID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	CustomerAddressMapping, err := lc.CustomerAddressMappingUsecase.FetchByID(c, CustomerAddressMappingID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: CustomerAddressMappingID})
		return
	}

	c.JSON(http.StatusOK, CustomerAddressMapping)
}

func (lc *CustomerAddressMappingController) Fetch(c *gin.Context) {

	CustomerAddressMapping, err := lc.CustomerAddressMappingUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CustomerAddressMapping)
}
