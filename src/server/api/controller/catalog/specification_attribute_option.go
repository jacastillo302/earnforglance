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

type SpecificationAttributeOptionController struct {
	SpecificationAttributeOptionUsecase domain.SpecificationAttributeOptionUsecase
	Env                                 *bootstrap.Env
}

func (tc *SpecificationAttributeOptionController) Create(c *gin.Context) {
	var task domain.SpecificationAttributeOption
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

	err = tc.SpecificationAttributeOptionUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SpecificationAttributeOption created successfully",
	})
}

func (tc *SpecificationAttributeOptionController) Update(c *gin.Context) {
	var task domain.SpecificationAttributeOption
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

	err = tc.SpecificationAttributeOptionUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SpecificationAttributeOption update successfully",
	})
}

func (tc *SpecificationAttributeOptionController) Delete(c *gin.Context) {
	var task domain.SpecificationAttributeOption
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

	err = tc.SpecificationAttributeOptionUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SpecificationAttributeOption update successfully",
	})
}

func (lc *SpecificationAttributeOptionController) FetchByID(c *gin.Context) {
	SpecificationAttributeOptionID := c.Query("id")
	if SpecificationAttributeOptionID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	SpecificationAttributeOption, err := lc.SpecificationAttributeOptionUsecase.FetchByID(c, SpecificationAttributeOptionID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: SpecificationAttributeOptionID})
		return
	}

	c.JSON(http.StatusOK, SpecificationAttributeOption)
}

func (lc *SpecificationAttributeOptionController) Fetch(c *gin.Context) {

	SpecificationAttributeOption, err := lc.SpecificationAttributeOptionUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SpecificationAttributeOption)
}
