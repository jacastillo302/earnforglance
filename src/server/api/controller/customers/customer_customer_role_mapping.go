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

type CustomerCustomerRoleMappingController struct {
	CustomerCustomerRoleMappingUsecase domain.CustomerCustomerRoleMappingUsecase
	Env                                *bootstrap.Env
}

func (tc *CustomerCustomerRoleMappingController) Create(c *gin.Context) {
	var task domain.CustomerCustomerRoleMapping
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

	err = tc.CustomerCustomerRoleMappingUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CustomerCustomerRoleMapping created successfully",
	})
}

func (tc *CustomerCustomerRoleMappingController) Update(c *gin.Context) {
	var task domain.CustomerCustomerRoleMapping
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

	err = tc.CustomerCustomerRoleMappingUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CustomerCustomerRoleMapping update successfully",
	})
}

func (tc *CustomerCustomerRoleMappingController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.CustomerCustomerRoleMappingUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *CustomerCustomerRoleMappingController) FetchByID(c *gin.Context) {
	CustomerCustomerRoleMappingID := c.Query("id")
	if CustomerCustomerRoleMappingID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	CustomerCustomerRoleMapping, err := lc.CustomerCustomerRoleMappingUsecase.FetchByID(c, CustomerCustomerRoleMappingID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: CustomerCustomerRoleMappingID})
		return
	}

	c.JSON(http.StatusOK, CustomerCustomerRoleMapping)
}

func (lc *CustomerCustomerRoleMappingController) Fetch(c *gin.Context) {

	CustomerCustomerRoleMapping, err := lc.CustomerCustomerRoleMappingUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CustomerCustomerRoleMapping)
}
