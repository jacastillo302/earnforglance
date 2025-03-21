package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/directory"

	"github.com/gin-gonic/gin"
)

type MeasureWeightController struct {
	MeasureWeightUsecase domain.MeasureWeightUsecase
	Env                  *bootstrap.Env
}

func (tc *MeasureWeightController) Create(c *gin.Context) {
	var task domain.MeasureWeight
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

	err = tc.MeasureWeightUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "MeasureWeight created successfully",
	})
}

func (tc *MeasureWeightController) Update(c *gin.Context) {
	var task domain.MeasureWeight
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

	err = tc.MeasureWeightUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "MeasureWeight update successfully",
	})
}

func (tc *MeasureWeightController) Delete(c *gin.Context) {
	var task domain.MeasureWeight
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

	err = tc.MeasureWeightUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "MeasureWeight update successfully",
	})
}

func (lc *MeasureWeightController) FetchByID(c *gin.Context) {
	MeasureWeightID := c.Query("id")
	if MeasureWeightID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	MeasureWeight, err := lc.MeasureWeightUsecase.FetchByID(c, MeasureWeightID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: MeasureWeightID})
		return
	}

	c.JSON(http.StatusOK, MeasureWeight)
}

func (lc *MeasureWeightController) Fetch(c *gin.Context) {

	MeasureWeight, err := lc.MeasureWeightUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, MeasureWeight)
}
