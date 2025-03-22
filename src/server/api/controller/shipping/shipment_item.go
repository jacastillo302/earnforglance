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

type ShipmentItemController struct {
	ShipmentItemUsecase domain.ShipmentItemUsecase
	Env                 *bootstrap.Env
}

func (tc *ShipmentItemController) Create(c *gin.Context) {
	var task domain.ShipmentItem
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

	err = tc.ShipmentItemUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ShipmentItem created successfully",
	})
}

func (tc *ShipmentItemController) Update(c *gin.Context) {
	var task domain.ShipmentItem
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

	err = tc.ShipmentItemUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ShipmentItem update successfully",
	})
}

func (tc *ShipmentItemController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ShipmentItemUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ShipmentItemController) FetchByID(c *gin.Context) {
	ShipmentItemID := c.Query("id")
	if ShipmentItemID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ShipmentItem, err := lc.ShipmentItemUsecase.FetchByID(c, ShipmentItemID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ShipmentItemID})
		return
	}

	c.JSON(http.StatusOK, ShipmentItem)
}

func (lc *ShipmentItemController) Fetch(c *gin.Context) {

	ShipmentItem, err := lc.ShipmentItemUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ShipmentItem)
}
