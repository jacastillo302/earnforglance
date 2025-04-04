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

type CustomerRoleController struct {
	CustomerRoleUsecase domain.CustomerRoleUsecase
	Env                 *bootstrap.Env
}

func (tc *CustomerRoleController) CreateMany(c *gin.Context) {
	var task []domain.CustomerRole
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

	err = tc.CustomerRoleUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CustomerRole created successfully",
	})
}

func (tc *CustomerRoleController) Create(c *gin.Context) {
	var task domain.CustomerRole
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

	err = tc.CustomerRoleUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CustomerRole created successfully",
	})
}

func (tc *CustomerRoleController) Update(c *gin.Context) {
	var task domain.CustomerRole
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

	err = tc.CustomerRoleUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CustomerRole update successfully",
	})
}

func (tc *CustomerRoleController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.CustomerRoleUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CustomerRole Record deleted successfully",
	})
}

func (lc *CustomerRoleController) FetchByID(c *gin.Context) {
	CustomerRoleID := c.Query("id")
	if CustomerRoleID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	CustomerRole, err := lc.CustomerRoleUsecase.FetchByID(c, CustomerRoleID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: CustomerRoleID})
		return
	}

	c.JSON(http.StatusOK, CustomerRole)
}

func (lc *CustomerRoleController) Fetch(c *gin.Context) {
	CustomerRole, err := lc.CustomerRoleUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CustomerRole)
}
