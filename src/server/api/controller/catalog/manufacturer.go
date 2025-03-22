package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/catalog"
	common "earnforglance/server/domain/common"

	"github.com/gin-gonic/gin"
)

type ManufacturerController struct {
	ManufacturerUsecase domain.ManufacturerUsecase
	Env                 *bootstrap.Env
}

func (tc *ManufacturerController) Create(c *gin.Context) {
	var task domain.Manufacturer
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

	err = tc.ManufacturerUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Manufacturer created successfully",
	})
}

func (tc *ManufacturerController) Update(c *gin.Context) {
	var task domain.Manufacturer
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

	err = tc.ManufacturerUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Manufacturer update successfully",
	})
}

func (tc *ManufacturerController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ManufacturerUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ManufacturerController) FetchByID(c *gin.Context) {
	ManufacturerID := c.Query("id")
	if ManufacturerID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	Manufacturer, err := lc.ManufacturerUsecase.FetchByID(c, ManufacturerID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ManufacturerID})
		return
	}

	c.JSON(http.StatusOK, Manufacturer)
}

func (lc *ManufacturerController) Fetch(c *gin.Context) {

	Manufacturer, err := lc.ManufacturerUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Manufacturer)
}
