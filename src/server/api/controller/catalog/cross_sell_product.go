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

type CrossSellProductController struct {
	CrossSellProductUsecase domain.CrossSellProductUsecase
	Env                     *bootstrap.Env
}

func (tc *CrossSellProductController) Create(c *gin.Context) {
	var task domain.CrossSellProduct
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

	err = tc.CrossSellProductUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CrossSellProduct created successfully",
	})
}

func (tc *CrossSellProductController) Update(c *gin.Context) {
	var task domain.CrossSellProduct
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

	err = tc.CrossSellProductUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CrossSellProduct update successfully",
	})
}

func (tc *CrossSellProductController) Delete(c *gin.Context) {
	var task domain.CrossSellProduct
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

	err = tc.CrossSellProductUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "CrossSellProduct update successfully",
	})
}

func (lc *CrossSellProductController) FetchByID(c *gin.Context) {
	CrossSellProductID := c.Query("id")
	if CrossSellProductID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	CrossSellProduct, err := lc.CrossSellProductUsecase.FetchByID(c, CrossSellProductID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: CrossSellProductID})
		return
	}

	c.JSON(http.StatusOK, CrossSellProduct)
}

func (lc *CrossSellProductController) Fetch(c *gin.Context) {

	CrossSellProduct, err := lc.CrossSellProductUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CrossSellProduct)
}
