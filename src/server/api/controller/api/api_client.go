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

type ApiClientController struct {
	ApiClientUsecase domain.ApiClientUsecase
	Env              *bootstrap.Env
}

func (tc *ApiClientController) CreateMany(c *gin.Context) {
	var task []domain.ApiClient
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

	err = tc.ApiClientUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ApiClient created successfully",
	})
}

func (tc *ApiClientController) Create(c *gin.Context) {
	var task domain.ApiClient
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

	err = tc.ApiClientUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ApiClient created successfully",
	})
}

func (tc *ApiClientController) Update(c *gin.Context) {
	var task domain.ApiClient
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

	err = tc.ApiClientUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ApiClient update successfully",
	})
}

func (tc *ApiClientController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ApiClientUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ApiClientController) FetchByID(c *gin.Context) {
	ApiClientID := c.Query("id")
	if ApiClientID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ApiClient, err := lc.ApiClientUsecase.FetchByID(c, ApiClientID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ApiClientID})
		return
	}

	c.JSON(http.StatusOK, ApiClient)
}

func (lc *ApiClientController) Fetch(c *gin.Context) {
	ApiClient, err := lc.ApiClientUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ApiClient)
}
