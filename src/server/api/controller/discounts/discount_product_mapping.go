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

type DiscountProductMappingController struct {
	DiscountProductMappingUsecase domain.DiscountProductMappingUsecase
	Env                           *bootstrap.Env
}

func (tc *DiscountProductMappingController) Create(c *gin.Context) {
	var task domain.DiscountProductMapping
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

	err = tc.DiscountProductMappingUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DiscountProductMapping created successfully",
	})
}

func (tc *DiscountProductMappingController) Update(c *gin.Context) {
	var task domain.DiscountProductMapping
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

	err = tc.DiscountProductMappingUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DiscountProductMapping update successfully",
	})
}

func (tc *DiscountProductMappingController) Delete(c *gin.Context) {
	var task domain.DiscountProductMapping
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

	err = tc.DiscountProductMappingUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DiscountProductMapping update successfully",
	})
}

func (lc *DiscountProductMappingController) FetchByID(c *gin.Context) {
	DiscountProductMappingID := c.Query("id")
	if DiscountProductMappingID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	DiscountProductMapping, err := lc.DiscountProductMappingUsecase.FetchByID(c, DiscountProductMappingID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: DiscountProductMappingID})
		return
	}

	c.JSON(http.StatusOK, DiscountProductMapping)
}

func (lc *DiscountProductMappingController) Fetch(c *gin.Context) {

	DiscountProductMapping, err := lc.DiscountProductMappingUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DiscountProductMapping)
}
