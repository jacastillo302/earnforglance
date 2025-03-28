package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/api"
	common "earnforglance/server/domain/common"

	"github.com/gin-gonic/gin"
)

type ApiSettingsController struct {
	ApiSettingsUsecase domain.ApiSettingsUsecase
	Env                *bootstrap.Env
}

func (tc *ApiSettingsController) CreateMany(c *gin.Context) {
	var task []domain.ApiSettings
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

	err = tc.ApiSettingsUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ApiSettings created successfully",
	})
}

func (tc *ApiSettingsController) Create(c *gin.Context) {
	var task domain.ApiSettings
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

	err = tc.ApiSettingsUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ApiSettings created successfully",
	})
}

func (tc *ApiSettingsController) Update(c *gin.Context) {
	var task domain.ApiSettings
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

	err = tc.ApiSettingsUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ApiSettings update successfully",
	})
}

func (tc *ApiSettingsController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ApiSettingsUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ApiSettingsController) FetchByID(c *gin.Context) {
	ApiSettingsID := c.Query("id")
	if ApiSettingsID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ApiSettings, err := lc.ApiSettingsUsecase.FetchByID(c, ApiSettingsID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ApiSettingsID})
		return
	}

	c.JSON(http.StatusOK, ApiSettings)
}

func (lc *ApiSettingsController) Fetch(c *gin.Context) {
	ApiSettings, err := lc.ApiSettingsUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ApiSettings)
}
