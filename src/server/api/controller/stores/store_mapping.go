package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/stores"

	"github.com/gin-gonic/gin"
)

type StoreMappingController struct {
	StoreMappingUsecase domain.StoreMappingUsecase
	Env                 *bootstrap.Env
}

func (tc *StoreMappingController) Create(c *gin.Context) {
	var task domain.StoreMapping
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

	err = tc.StoreMappingUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "StoreMapping created successfully",
	})
}

func (tc *StoreMappingController) Update(c *gin.Context) {
	var task domain.StoreMapping
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

	err = tc.StoreMappingUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "StoreMapping update successfully",
	})
}

func (tc *StoreMappingController) Delete(c *gin.Context) {
	var task domain.StoreMapping
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

	err = tc.StoreMappingUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "StoreMapping update successfully",
	})
}

func (lc *StoreMappingController) FetchByID(c *gin.Context) {
	StoreMappingID := c.Query("id")
	if StoreMappingID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	StoreMapping, err := lc.StoreMappingUsecase.FetchByID(c, StoreMappingID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: StoreMappingID})
		return
	}

	c.JSON(http.StatusOK, StoreMapping)
}

func (lc *StoreMappingController) Fetch(c *gin.Context) {

	StoreMapping, err := lc.StoreMappingUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, StoreMapping)
}
