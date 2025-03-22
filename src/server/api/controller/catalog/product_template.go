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

type ProductTemplateController struct {
	ProductTemplateUsecase domain.ProductTemplateUsecase
	Env                    *bootstrap.Env
}

func (tc *ProductTemplateController) Create(c *gin.Context) {
	var task domain.ProductTemplate
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

	err = tc.ProductTemplateUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductTemplate created successfully",
	})
}

func (tc *ProductTemplateController) Update(c *gin.Context) {
	var task domain.ProductTemplate
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

	err = tc.ProductTemplateUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductTemplate update successfully",
	})
}

func (tc *ProductTemplateController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ProductTemplateUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ProductTemplateController) FetchByID(c *gin.Context) {
	ProductTemplateID := c.Query("id")
	if ProductTemplateID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ProductTemplate, err := lc.ProductTemplateUsecase.FetchByID(c, ProductTemplateID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ProductTemplateID})
		return
	}

	c.JSON(http.StatusOK, ProductTemplate)
}

func (lc *ProductTemplateController) Fetch(c *gin.Context) {

	ProductTemplate, err := lc.ProductTemplateUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProductTemplate)
}
