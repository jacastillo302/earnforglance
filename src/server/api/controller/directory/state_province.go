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

type StateProvinceController struct {
	StateProvinceUsecase domain.StateProvinceUsecase
	Env                  *bootstrap.Env
}

func (tc *StateProvinceController) Create(c *gin.Context) {
	var task domain.StateProvince
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

	err = tc.StateProvinceUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "StateProvince created successfully",
	})
}

func (tc *StateProvinceController) Update(c *gin.Context) {
	var task domain.StateProvince
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

	err = tc.StateProvinceUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "StateProvince update successfully",
	})
}

func (tc *StateProvinceController) Delete(c *gin.Context) {
	var task domain.StateProvince
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

	err = tc.StateProvinceUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "StateProvince update successfully",
	})
}

func (lc *StateProvinceController) FetchByID(c *gin.Context) {
	StateProvinceID := c.Query("id")
	if StateProvinceID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	StateProvince, err := lc.StateProvinceUsecase.FetchByID(c, StateProvinceID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: StateProvinceID})
		return
	}

	c.JSON(http.StatusOK, StateProvince)
}

func (lc *StateProvinceController) Fetch(c *gin.Context) {

	StateProvince, err := lc.StateProvinceUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, StateProvince)
}
