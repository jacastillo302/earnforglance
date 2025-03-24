package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/security"

	"github.com/gin-gonic/gin"
)

type PermissionRecordCustomerRoleMappingController struct {
	PermissionRecordCustomerRoleMappingUsecase domain.PermissionRecordCustomerRoleMappingUsecase
	Env                                        *bootstrap.Env
}

func (tc *PermissionRecordCustomerRoleMappingController) CreateMany(c *gin.Context) {
	var task []domain.PermissionRecordCustomerRoleMapping
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

	err = tc.PermissionRecordCustomerRoleMappingUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PermissionRecordCustomerRoleMapping created successfully",
	})
}

func (tc *PermissionRecordCustomerRoleMappingController) Create(c *gin.Context) {
	var task domain.PermissionRecordCustomerRoleMapping
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

	err = tc.PermissionRecordCustomerRoleMappingUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PermissionRecordCustomerRoleMapping created successfully",
	})
}

func (tc *PermissionRecordCustomerRoleMappingController) Update(c *gin.Context) {
	var task domain.PermissionRecordCustomerRoleMapping
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

	err = tc.PermissionRecordCustomerRoleMappingUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PermissionRecordCustomerRoleMapping update successfully",
	})
}

func (tc *PermissionRecordCustomerRoleMappingController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.PermissionRecordCustomerRoleMappingUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *PermissionRecordCustomerRoleMappingController) FetchByID(c *gin.Context) {
	PermissionRecordCustomerRoleMappingID := c.Query("id")
	if PermissionRecordCustomerRoleMappingID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	PermissionRecordCustomerRoleMapping, err := lc.PermissionRecordCustomerRoleMappingUsecase.FetchByID(c, PermissionRecordCustomerRoleMappingID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: PermissionRecordCustomerRoleMappingID})
		return
	}

	c.JSON(http.StatusOK, PermissionRecordCustomerRoleMapping)
}

func (lc *PermissionRecordCustomerRoleMappingController) Fetch(c *gin.Context) {

	PermissionRecordCustomerRoleMapping, err := lc.PermissionRecordCustomerRoleMappingUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, PermissionRecordCustomerRoleMapping)
}
