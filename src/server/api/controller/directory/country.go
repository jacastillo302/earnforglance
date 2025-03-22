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

type CountryController struct {
	CountryUsecase domain.CountryUsecase
	Env            *bootstrap.Env
}

func (tc *CountryController) Create(c *gin.Context) {
	var task domain.Country
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

	err = tc.CountryUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Country created successfully",
	})
}

func (tc *CountryController) Update(c *gin.Context) {
	var task domain.Country
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

	err = tc.CountryUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Country update successfully",
	})
}

func (tc *CountryController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.CountryUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *CountryController) FetchByID(c *gin.Context) {
	CountryID := c.Query("id")
	if CountryID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	Country, err := lc.CountryUsecase.FetchByID(c, CountryID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: CountryID})
		return
	}

	c.JSON(http.StatusOK, Country)
}

func (lc *CountryController) Fetch(c *gin.Context) {

	Country, err := lc.CountryUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Country)
}
