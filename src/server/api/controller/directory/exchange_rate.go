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

type ExchangeRateController struct {
	ExchangeRateUsecase domain.ExchangeRateUsecase
	Env                 *bootstrap.Env
}

func (tc *ExchangeRateController) Create(c *gin.Context) {
	var task domain.ExchangeRate
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

	err = tc.ExchangeRateUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ExchangeRate created successfully",
	})
}

func (tc *ExchangeRateController) Update(c *gin.Context) {
	var task domain.ExchangeRate
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

	err = tc.ExchangeRateUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ExchangeRate update successfully",
	})
}

func (tc *ExchangeRateController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ExchangeRateUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ExchangeRateController) FetchByID(c *gin.Context) {
	ExchangeRateID := c.Query("id")
	if ExchangeRateID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ExchangeRate, err := lc.ExchangeRateUsecase.FetchByID(c, ExchangeRateID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ExchangeRateID})
		return
	}

	c.JSON(http.StatusOK, ExchangeRate)
}

func (lc *ExchangeRateController) Fetch(c *gin.Context) {

	ExchangeRate, err := lc.ExchangeRateUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ExchangeRate)
}
