package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/directory"

	"github.com/gin-gonic/gin"
)

type CurrencyController struct {
	CurrencyUsecase domain.CurrencyUsecase
	Env             *bootstrap.Env
}

func (tc *CurrencyController) Create(c *gin.Context) {
	var task domain.Currency
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

	err = tc.CurrencyUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Currency created successfully",
	})
}

func (tc *CurrencyController) Update(c *gin.Context) {
	var task domain.Currency
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

	err = tc.CurrencyUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Currency update successfully",
	})
}

func (tc *CurrencyController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.CurrencyUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *CurrencyController) FetchByID(c *gin.Context) {
	CurrencyID := c.Query("id")
	if CurrencyID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	Currency, err := lc.CurrencyUsecase.FetchByID(c, CurrencyID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: CurrencyID})
		return
	}

	c.JSON(http.StatusOK, Currency)
}

func (lc *CurrencyController) Fetch(c *gin.Context) {

	Currency, err := lc.CurrencyUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Currency)
}
