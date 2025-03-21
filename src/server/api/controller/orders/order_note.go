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

type OrderNoteController struct {
	OrderNoteUsecase domain.OrderNoteUsecase
	Env              *bootstrap.Env
}

func (tc *OrderNoteController) Create(c *gin.Context) {
	var task domain.OrderNote
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

	err = tc.OrderNoteUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "OrderNote created successfully",
	})
}

func (tc *OrderNoteController) Update(c *gin.Context) {
	var task domain.OrderNote
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

	err = tc.OrderNoteUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "OrderNote update successfully",
	})
}

func (tc *OrderNoteController) Delete(c *gin.Context) {
	var task domain.OrderNote
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

	err = tc.OrderNoteUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "OrderNote update successfully",
	})
}

func (lc *OrderNoteController) FetchByID(c *gin.Context) {
	OrderNoteID := c.Query("id")
	if OrderNoteID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	OrderNote, err := lc.OrderNoteUsecase.FetchByID(c, OrderNoteID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: OrderNoteID})
		return
	}

	c.JSON(http.StatusOK, OrderNote)
}

func (lc *OrderNoteController) Fetch(c *gin.Context) {

	OrderNote, err := lc.OrderNoteUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, OrderNote)
}
