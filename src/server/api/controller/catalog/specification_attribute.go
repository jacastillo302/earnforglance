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

type SpecificationAttributeController struct {
	SpecificationAttributeUsecase domain.SpecificationAttributeUsecase
	Env                           *bootstrap.Env
}

func (tc *SpecificationAttributeController) Create(c *gin.Context) {
	var task domain.SpecificationAttribute
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

	err = tc.SpecificationAttributeUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SpecificationAttribute created successfully",
	})
}

func (tc *SpecificationAttributeController) Update(c *gin.Context) {
	var task domain.SpecificationAttribute
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

	err = tc.SpecificationAttributeUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SpecificationAttribute update successfully",
	})
}

func (tc *SpecificationAttributeController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.SpecificationAttributeUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *SpecificationAttributeController) FetchByID(c *gin.Context) {
	SpecificationAttributeID := c.Query("id")
	if SpecificationAttributeID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	SpecificationAttribute, err := lc.SpecificationAttributeUsecase.FetchByID(c, SpecificationAttributeID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: SpecificationAttributeID})
		return
	}

	c.JSON(http.StatusOK, SpecificationAttribute)
}

func (lc *SpecificationAttributeController) Fetch(c *gin.Context) {

	SpecificationAttribute, err := lc.SpecificationAttributeUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SpecificationAttribute)
}
