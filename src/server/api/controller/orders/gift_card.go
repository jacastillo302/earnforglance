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

type GiftCardController struct {
	GiftCardUsecase domain.GiftCardUsecase
	Env             *bootstrap.Env
}

func (tc *GiftCardController) Create(c *gin.Context) {
	var task domain.GiftCard
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

	err = tc.GiftCardUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "GiftCard created successfully",
	})
}

func (tc *GiftCardController) Update(c *gin.Context) {
	var task domain.GiftCard
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

	err = tc.GiftCardUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "GiftCard update successfully",
	})
}

func (tc *GiftCardController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.GiftCardUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *GiftCardController) FetchByID(c *gin.Context) {
	GiftCardID := c.Query("id")
	if GiftCardID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	GiftCard, err := lc.GiftCardUsecase.FetchByID(c, GiftCardID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: GiftCardID})
		return
	}

	c.JSON(http.StatusOK, GiftCard)
}

func (lc *GiftCardController) Fetch(c *gin.Context) {

	GiftCard, err := lc.GiftCardUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GiftCard)
}
