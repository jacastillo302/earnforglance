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

type ProductAttributeCombinationController struct {
	ProductAttributeCombinationUsecase domain.ProductAttributeCombinationUsecase
	Env                                *bootstrap.Env
}

func (tc *ProductAttributeCombinationController) Create(c *gin.Context) {
	var task domain.ProductAttributeCombination
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

	err = tc.ProductAttributeCombinationUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductAttributeCombination created successfully",
	})
}

func (tc *ProductAttributeCombinationController) Update(c *gin.Context) {
	var task domain.ProductAttributeCombination
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

	err = tc.ProductAttributeCombinationUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductAttributeCombination update successfully",
	})
}

func (tc *ProductAttributeCombinationController) Delete(c *gin.Context) {
	var task domain.ProductAttributeCombination
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

	err = tc.ProductAttributeCombinationUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductAttributeCombination update successfully",
	})
}

func (lc *ProductAttributeCombinationController) FetchByID(c *gin.Context) {
	ProductAttributeCombinationID := c.Query("id")
	if ProductAttributeCombinationID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ProductAttributeCombination, err := lc.ProductAttributeCombinationUsecase.FetchByID(c, ProductAttributeCombinationID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ProductAttributeCombinationID})
		return
	}

	c.JSON(http.StatusOK, ProductAttributeCombination)
}

func (lc *ProductAttributeCombinationController) Fetch(c *gin.Context) {

	ProductAttributeCombination, err := lc.ProductAttributeCombinationUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProductAttributeCombination)
}
