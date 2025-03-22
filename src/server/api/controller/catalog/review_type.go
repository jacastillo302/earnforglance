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

type ReviewTypeController struct {
	ReviewTypeUsecase domain.ReviewTypeUsecase
	Env               *bootstrap.Env
}

func (tc *ReviewTypeController) Create(c *gin.Context) {
	var task domain.ReviewType
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

	err = tc.ReviewTypeUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ReviewType created successfully",
	})
}

func (tc *ReviewTypeController) Update(c *gin.Context) {
	var task domain.ReviewType
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

	err = tc.ReviewTypeUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ReviewType update successfully",
	})
}

func (tc *ReviewTypeController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ReviewTypeUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ReviewTypeController) FetchByID(c *gin.Context) {
	ReviewTypeID := c.Query("id")
	if ReviewTypeID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ReviewType, err := lc.ReviewTypeUsecase.FetchByID(c, ReviewTypeID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ReviewTypeID})
		return
	}

	c.JSON(http.StatusOK, ReviewType)
}

func (lc *ReviewTypeController) Fetch(c *gin.Context) {

	ReviewType, err := lc.ReviewTypeUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ReviewType)
}
