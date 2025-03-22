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

type ManufacturerTemplateController struct {
	ManufacturerTemplateUsecase domain.ManufacturerTemplateUsecase
	Env                         *bootstrap.Env
}

func (tc *ManufacturerTemplateController) Create(c *gin.Context) {
	var task domain.ManufacturerTemplate
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

	err = tc.ManufacturerTemplateUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ManufacturerTemplate created successfully",
	})
}

func (tc *ManufacturerTemplateController) Update(c *gin.Context) {
	var task domain.ManufacturerTemplate
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

	err = tc.ManufacturerTemplateUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ManufacturerTemplate update successfully",
	})
}

func (tc *ManufacturerTemplateController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ManufacturerTemplateUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ManufacturerTemplateController) FetchByID(c *gin.Context) {
	ManufacturerTemplateID := c.Query("id")
	if ManufacturerTemplateID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ManufacturerTemplate, err := lc.ManufacturerTemplateUsecase.FetchByID(c, ManufacturerTemplateID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ManufacturerTemplateID})
		return
	}

	c.JSON(http.StatusOK, ManufacturerTemplate)
}

func (lc *ManufacturerTemplateController) Fetch(c *gin.Context) {

	ManufacturerTemplate, err := lc.ManufacturerTemplateUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ManufacturerTemplate)
}
