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

type RewardPointsHistoryController struct {
	RewardPointsHistoryUsecase domain.RewardPointsHistoryUsecase
	Env                        *bootstrap.Env
}

func (tc *RewardPointsHistoryController) Create(c *gin.Context) {
	var task domain.RewardPointsHistory
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

	err = tc.RewardPointsHistoryUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "RewardPointsHistory created successfully",
	})
}

func (tc *RewardPointsHistoryController) Update(c *gin.Context) {
	var task domain.RewardPointsHistory
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

	err = tc.RewardPointsHistoryUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "RewardPointsHistory update successfully",
	})
}

func (tc *RewardPointsHistoryController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.RewardPointsHistoryUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *RewardPointsHistoryController) FetchByID(c *gin.Context) {
	RewardPointsHistoryID := c.Query("id")
	if RewardPointsHistoryID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	RewardPointsHistory, err := lc.RewardPointsHistoryUsecase.FetchByID(c, RewardPointsHistoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: RewardPointsHistoryID})
		return
	}

	c.JSON(http.StatusOK, RewardPointsHistory)
}

func (lc *RewardPointsHistoryController) Fetch(c *gin.Context) {

	RewardPointsHistory, err := lc.RewardPointsHistoryUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, RewardPointsHistory)
}
