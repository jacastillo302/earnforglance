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

type BackInStockSubscriptionController struct {
	BackInStockSubscriptionUsecase domain.BackInStockSubscriptionUsecase
	Env                            *bootstrap.Env
}

func (tc *BackInStockSubscriptionController) Create(c *gin.Context) {
	var task domain.BackInStockSubscription
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

	err = tc.BackInStockSubscriptionUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "BackInStockSubscription created successfully",
	})
}

func (tc *BackInStockSubscriptionController) Update(c *gin.Context) {
	var task domain.BackInStockSubscription
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

	err = tc.BackInStockSubscriptionUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "BackInStockSubscription update successfully",
	})
}

func (tc *BackInStockSubscriptionController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.BackInStockSubscriptionUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *BackInStockSubscriptionController) FetchByID(c *gin.Context) {
	BackInStockSubscriptionID := c.Query("id")
	if BackInStockSubscriptionID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	BackInStockSubscription, err := lc.BackInStockSubscriptionUsecase.FetchByID(c, BackInStockSubscriptionID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: BackInStockSubscriptionID})
		return
	}

	c.JSON(http.StatusOK, BackInStockSubscription)
}

func (lc *BackInStockSubscriptionController) Fetch(c *gin.Context) {

	BackInStockSubscription, err := lc.BackInStockSubscriptionUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, BackInStockSubscription)
}
