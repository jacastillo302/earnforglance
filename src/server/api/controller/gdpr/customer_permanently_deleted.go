package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/gdpr"

	"github.com/gin-gonic/gin"
)

type CustomerPermanentlyDeletedController struct {
	CustomerPermanentlyDeletedUsecase domain.CustomerPermanentlyDeletedUsecase
	Env                               *bootstrap.Env
}

func (tc *CustomerPermanentlyDeletedController) Create(c *gin.Context) {
	var task domain.CustomerPermanentlyDeleted
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

	err = tc.CustomerPermanentlyDeletedUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CustomerPermanentlyDeleted created successfully",
	})
}

func (tc *CustomerPermanentlyDeletedController) Update(c *gin.Context) {
	var task domain.CustomerPermanentlyDeleted
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

	err = tc.CustomerPermanentlyDeletedUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CustomerPermanentlyDeleted update successfully",
	})
}

func (tc *CustomerPermanentlyDeletedController) Delete(c *gin.Context) {
	var task domain.CustomerPermanentlyDeleted
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

	err = tc.CustomerPermanentlyDeletedUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CustomerPermanentlyDeleted update successfully",
	})
}

func (lc *CustomerPermanentlyDeletedController) FetchByID(c *gin.Context) {
	CustomerPermanentlyDeletedID := c.Query("id")
	if CustomerPermanentlyDeletedID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	CustomerPermanentlyDeleted, err := lc.CustomerPermanentlyDeletedUsecase.FetchByID(c, CustomerPermanentlyDeletedID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: CustomerPermanentlyDeletedID})
		return
	}

	c.JSON(http.StatusOK, CustomerPermanentlyDeleted)
}

func (lc *CustomerPermanentlyDeletedController) Fetch(c *gin.Context) {

	CustomerPermanentlyDeleted, err := lc.CustomerPermanentlyDeletedUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CustomerPermanentlyDeleted)
}
