package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/discounts"

	"github.com/gin-gonic/gin"
)

type DiscountCategoryMappingController struct {
	DiscountCategoryMappingUsecase domain.DiscountCategoryMappingUsecase
	Env                            *bootstrap.Env
}

func (tc *DiscountCategoryMappingController) Create(c *gin.Context) {
	var task domain.DiscountCategoryMapping
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

	err = tc.DiscountCategoryMappingUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DiscountCategoryMapping created successfully",
	})
}

func (tc *DiscountCategoryMappingController) Update(c *gin.Context) {
	var task domain.DiscountCategoryMapping
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

	err = tc.DiscountCategoryMappingUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DiscountCategoryMapping update successfully",
	})
}

func (tc *DiscountCategoryMappingController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.DiscountCategoryMappingUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *DiscountCategoryMappingController) FetchByID(c *gin.Context) {
	DiscountCategoryMappingID := c.Query("id")
	if DiscountCategoryMappingID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	DiscountCategoryMapping, err := lc.DiscountCategoryMappingUsecase.FetchByID(c, DiscountCategoryMappingID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: DiscountCategoryMappingID})
		return
	}

	c.JSON(http.StatusOK, DiscountCategoryMapping)
}

func (lc *DiscountCategoryMappingController) Fetch(c *gin.Context) {

	DiscountCategoryMapping, err := lc.DiscountCategoryMappingUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DiscountCategoryMapping)
}
