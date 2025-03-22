package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/orders"

	"github.com/gin-gonic/gin"
)

type ShoppingCartItemController struct {
	ShoppingCartItemUsecase domain.ShoppingCartItemUsecase
	Env                     *bootstrap.Env
}

func (tc *ShoppingCartItemController) Create(c *gin.Context) {
	var task domain.ShoppingCartItem
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

	err = tc.ShoppingCartItemUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ShoppingCartItem created successfully",
	})
}

func (tc *ShoppingCartItemController) Update(c *gin.Context) {
	var task domain.ShoppingCartItem
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

	err = tc.ShoppingCartItemUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ShoppingCartItem update successfully",
	})
}

func (tc *ShoppingCartItemController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ShoppingCartItemUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ShoppingCartItemController) FetchByID(c *gin.Context) {
	ShoppingCartItemID := c.Query("id")
	if ShoppingCartItemID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ShoppingCartItem, err := lc.ShoppingCartItemUsecase.FetchByID(c, ShoppingCartItemID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ShoppingCartItemID})
		return
	}

	c.JSON(http.StatusOK, ShoppingCartItem)
}

func (lc *ShoppingCartItemController) Fetch(c *gin.Context) {

	ShoppingCartItem, err := lc.ShoppingCartItemUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ShoppingCartItem)
}
