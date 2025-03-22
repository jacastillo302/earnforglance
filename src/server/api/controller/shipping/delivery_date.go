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

type DeliveryDateController struct {
	DeliveryDateUsecase domain.DeliveryDateUsecase
	Env                 *bootstrap.Env
}

func (tc *DeliveryDateController) Create(c *gin.Context) {
	var task domain.DeliveryDate
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

	err = tc.DeliveryDateUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DeliveryDate created successfully",
	})
}

func (tc *DeliveryDateController) Update(c *gin.Context) {
	var task domain.DeliveryDate
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

	err = tc.DeliveryDateUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DeliveryDate update successfully",
	})
}

func (tc *DeliveryDateController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.DeliveryDateUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *DeliveryDateController) FetchByID(c *gin.Context) {
	DeliveryDateID := c.Query("id")
	if DeliveryDateID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	DeliveryDate, err := lc.DeliveryDateUsecase.FetchByID(c, DeliveryDateID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: DeliveryDateID})
		return
	}

	c.JSON(http.StatusOK, DeliveryDate)
}

func (lc *DeliveryDateController) Fetch(c *gin.Context) {

	DeliveryDate, err := lc.DeliveryDateUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DeliveryDate)
}
