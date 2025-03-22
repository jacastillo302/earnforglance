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

type RelatedProductController struct {
	RelatedProductUsecase domain.RelatedProductUsecase
	Env                   *bootstrap.Env
}

func (tc *RelatedProductController) Create(c *gin.Context) {
	var task domain.RelatedProduct
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

	err = tc.RelatedProductUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "RelatedProduct created successfully",
	})
}

func (tc *RelatedProductController) Update(c *gin.Context) {
	var task domain.RelatedProduct
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

	err = tc.RelatedProductUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "RelatedProduct update successfully",
	})
}

func (tc *RelatedProductController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.RelatedProductUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *RelatedProductController) FetchByID(c *gin.Context) {
	RelatedProductID := c.Query("id")
	if RelatedProductID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	RelatedProduct, err := lc.RelatedProductUsecase.FetchByID(c, RelatedProductID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: RelatedProductID})
		return
	}

	c.JSON(http.StatusOK, RelatedProduct)
}

func (lc *RelatedProductController) Fetch(c *gin.Context) {

	RelatedProduct, err := lc.RelatedProductUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, RelatedProduct)
}
