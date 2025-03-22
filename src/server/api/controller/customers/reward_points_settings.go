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

type RewardPointsSettingsController struct {
	RewardPointsSettingsUsecase domain.RewardPointsSettingsUsecase
	Env                         *bootstrap.Env
}

func (tc *RewardPointsSettingsController) Create(c *gin.Context) {
	var task domain.RewardPointsSettings
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

	err = tc.RewardPointsSettingsUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "RewardPointsSettings created successfully",
	})
}

func (tc *RewardPointsSettingsController) Update(c *gin.Context) {
	var task domain.RewardPointsSettings
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

	err = tc.RewardPointsSettingsUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "RewardPointsSettings update successfully",
	})
}

func (tc *RewardPointsSettingsController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.RewardPointsSettingsUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *RewardPointsSettingsController) FetchByID(c *gin.Context) {
	RewardPointsSettingsID := c.Query("id")
	if RewardPointsSettingsID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	RewardPointsSettings, err := lc.RewardPointsSettingsUsecase.FetchByID(c, RewardPointsSettingsID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: RewardPointsSettingsID})
		return
	}

	c.JSON(http.StatusOK, RewardPointsSettings)
}

func (lc *RewardPointsSettingsController) Fetch(c *gin.Context) {

	RewardPointsSettings, err := lc.RewardPointsSettingsUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, RewardPointsSettings)
}
