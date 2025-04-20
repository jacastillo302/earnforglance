package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/public"

	"github.com/gin-gonic/gin"
)

type CatalogController struct {
	CatalogUsecase domain.CatalogtUsecase
	Env            *bootstrap.Env
}

func (cc *CatalogController) GetProducts(c *gin.Context) {
	var request domain.ProductRequest

	filter, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Failed to read request body"})
		return
	}

	err = json.Unmarshal(filter, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Invalid request body"})
		return
	}

	productResponse, err := cc.CatalogUsecase.GetProducts(c, request)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, productResponse)
}

func (cc *CatalogController) GetCategories(c *gin.Context) {
	var request domain.CategoryRequest

	filter, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Failed to read request body"})
		return
	}

	err = json.Unmarshal(filter, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Invalid request body"})
		return
	}

	productResponse, err := cc.CatalogUsecase.GetCategories(c, request)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, productResponse)
}
