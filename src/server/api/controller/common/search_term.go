package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"

	"github.com/gin-gonic/gin"
)

type SearchTermController struct {
	SearchTermUsecase common.SearchTermUsecase
	Env               *bootstrap.Env
}

func (tc *SearchTermController) Create(c *gin.Context) {
	var task common.SearchTerm
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

	err = tc.SearchTermUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SearchTerm created successfully",
	})
}

func (tc *SearchTermController) Update(c *gin.Context) {
	var task common.SearchTerm
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

	err = tc.SearchTermUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SearchTerm update successfully",
	})
}

func (tc *SearchTermController) Delete(c *gin.Context) {
	var task common.SearchTerm
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

	err = tc.SearchTermUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "SearchTerm update successfully",
	})
}

func (lc *SearchTermController) FetchByID(c *gin.Context) {
	SearchTermID := c.Query("id")
	if SearchTermID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	SearchTerm, err := lc.SearchTermUsecase.FetchByID(c, SearchTermID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: SearchTermID})
		return
	}

	c.JSON(http.StatusOK, SearchTerm)
}

func (lc *SearchTermController) Fetch(c *gin.Context) {

	SearchTerm, err := lc.SearchTermUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SearchTerm)
}
