package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/tax"

	"github.com/gin-gonic/gin"
)

type TaxCategoryController struct {
	TaxCategoryUsecase domain.TaxCategoryUsecase
	Env                *bootstrap.Env
}

func (tc *TaxCategoryController) Create(c *gin.Context) {
	var task domain.TaxCategory
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

	err = tc.TaxCategoryUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "TaxCategory created successfully",
	})
}

func (tc *TaxCategoryController) Update(c *gin.Context) {
	var task domain.TaxCategory
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

	err = tc.TaxCategoryUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "TaxCategory update successfully",
	})
}

func (tc *TaxCategoryController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.TaxCategoryUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *TaxCategoryController) FetchByID(c *gin.Context) {
	TaxCategoryID := c.Query("id")
	if TaxCategoryID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	TaxCategory, err := lc.TaxCategoryUsecase.FetchByID(c, TaxCategoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: TaxCategoryID})
		return
	}

	c.JSON(http.StatusOK, TaxCategory)
}

func (lc *TaxCategoryController) Fetch(c *gin.Context) {

	TaxCategory, err := lc.TaxCategoryUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, TaxCategory)
}
