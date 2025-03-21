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

type PickupPointController struct {
	PickupPointUsecase domain.PickupPointUsecase
	Env                *bootstrap.Env
}

func (tc *PickupPointController) Create(c *gin.Context) {
	var task domain.PickupPoint
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

	err = tc.PickupPointUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PickupPoint created successfully",
	})
}

func (tc *PickupPointController) Update(c *gin.Context) {
	var task domain.PickupPoint
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

	err = tc.PickupPointUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PickupPoint update successfully",
	})
}

func (tc *PickupPointController) Delete(c *gin.Context) {
	var task domain.PickupPoint
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

	err = tc.PickupPointUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PickupPoint update successfully",
	})
}

func (lc *PickupPointController) FetchByID(c *gin.Context) {
	PickupPointID := c.Query("id")
	if PickupPointID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	PickupPoint, err := lc.PickupPointUsecase.FetchByID(c, PickupPointID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: PickupPointID})
		return
	}

	c.JSON(http.StatusOK, PickupPoint)
}

func (lc *PickupPointController) Fetch(c *gin.Context) {

	PickupPoint, err := lc.PickupPointUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, PickupPoint)
}
