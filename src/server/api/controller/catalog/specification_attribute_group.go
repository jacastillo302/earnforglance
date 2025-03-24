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

type SpecificationAttributeGroupController struct {
	SpecificationAttributeGroupUsecase domain.SpecificationAttributeGroupUsecase
	Env                                *bootstrap.Env
}

func (tc *SpecificationAttributeGroupController) CreateMany(c *gin.Context) {
	var task []domain.SpecificationAttributeGroup
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

	err = tc.SpecificationAttributeGroupUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SpecificationAttributeGroup created successfully",
	})
}

func (tc *SpecificationAttributeGroupController) Create(c *gin.Context) {
	var task domain.SpecificationAttributeGroup
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

	err = tc.SpecificationAttributeGroupUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SpecificationAttributeGroup created successfully",
	})
}

func (tc *SpecificationAttributeGroupController) Update(c *gin.Context) {
	var task domain.SpecificationAttributeGroup
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

	err = tc.SpecificationAttributeGroupUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SpecificationAttributeGroup update successfully",
	})
}

func (tc *SpecificationAttributeGroupController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.SpecificationAttributeGroupUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *SpecificationAttributeGroupController) FetchByID(c *gin.Context) {
	SpecificationAttributeGroupID := c.Query("id")
	if SpecificationAttributeGroupID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	SpecificationAttributeGroup, err := lc.SpecificationAttributeGroupUsecase.FetchByID(c, SpecificationAttributeGroupID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: SpecificationAttributeGroupID})
		return
	}

	c.JSON(http.StatusOK, SpecificationAttributeGroup)
}

func (lc *SpecificationAttributeGroupController) Fetch(c *gin.Context) {

	SpecificationAttributeGroup, err := lc.SpecificationAttributeGroupUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SpecificationAttributeGroup)
}
