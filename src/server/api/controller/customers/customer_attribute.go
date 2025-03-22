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

type CustomerAttributeController struct {
	CustomerAttributeUsecase domain.CustomerAttributeUsecase
	Env                      *bootstrap.Env
}

func (tc *CustomerAttributeController) Create(c *gin.Context) {
	var task domain.CustomerAttribute
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

	err = tc.CustomerAttributeUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CustomerAttribute created successfully",
	})
}

func (tc *CustomerAttributeController) Update(c *gin.Context) {
	var task domain.CustomerAttribute
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

	err = tc.CustomerAttributeUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CustomerAttribute update successfully",
	})
}

func (tc *CustomerAttributeController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.CustomerAttributeUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *CustomerAttributeController) FetchByID(c *gin.Context) {
	CustomerAttributeID := c.Query("id")
	if CustomerAttributeID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	CustomerAttribute, err := lc.CustomerAttributeUsecase.FetchByID(c, CustomerAttributeID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: CustomerAttributeID})
		return
	}

	c.JSON(http.StatusOK, CustomerAttribute)
}

func (lc *CustomerAttributeController) Fetch(c *gin.Context) {

	CustomerAttribute, err := lc.CustomerAttributeUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CustomerAttribute)
}
