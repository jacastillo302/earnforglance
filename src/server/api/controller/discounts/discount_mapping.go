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

type DiscountMappingController struct {
	DiscountMappingUsecase domain.DiscountMappingUsecase
	Env                    *bootstrap.Env
}

func (tc *DiscountMappingController) Create(c *gin.Context) {
	var task domain.DiscountMapping
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

	err = tc.DiscountMappingUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DiscountMapping created successfully",
	})
}

func (tc *DiscountMappingController) Update(c *gin.Context) {
	var task domain.DiscountMapping
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

	err = tc.DiscountMappingUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DiscountMapping update successfully",
	})
}

func (tc *DiscountMappingController) Delete(c *gin.Context) {
	var task domain.DiscountMapping
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

	err = tc.DiscountMappingUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "DiscountMapping update successfully",
	})
}

func (lc *DiscountMappingController) FetchByID(c *gin.Context) {
	DiscountMappingID := c.Query("id")
	if DiscountMappingID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	DiscountMapping, err := lc.DiscountMappingUsecase.FetchByID(c, DiscountMappingID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: DiscountMappingID})
		return
	}

	c.JSON(http.StatusOK, DiscountMapping)
}

func (lc *DiscountMappingController) Fetch(c *gin.Context) {

	DiscountMapping, err := lc.DiscountMappingUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DiscountMapping)
}
