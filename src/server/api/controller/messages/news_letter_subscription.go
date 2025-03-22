package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/messages"

	"github.com/gin-gonic/gin"
)

type NewsLetterSubscriptionController struct {
	NewsLetterSubscriptionUsecase domain.NewsLetterSubscriptionUsecase
	Env                           *bootstrap.Env
}

func (tc *NewsLetterSubscriptionController) Create(c *gin.Context) {
	var task domain.NewsLetterSubscription
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

	err = tc.NewsLetterSubscriptionUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "NewsLetterSubscription created successfully",
	})
}

func (tc *NewsLetterSubscriptionController) Update(c *gin.Context) {
	var task domain.NewsLetterSubscription
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

	err = tc.NewsLetterSubscriptionUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "NewsLetterSubscription update successfully",
	})
}

func (tc *NewsLetterSubscriptionController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.NewsLetterSubscriptionUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *NewsLetterSubscriptionController) FetchByID(c *gin.Context) {
	NewsLetterSubscriptionID := c.Query("id")
	if NewsLetterSubscriptionID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	NewsLetterSubscription, err := lc.NewsLetterSubscriptionUsecase.FetchByID(c, NewsLetterSubscriptionID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: NewsLetterSubscriptionID})
		return
	}

	c.JSON(http.StatusOK, NewsLetterSubscription)
}

func (lc *NewsLetterSubscriptionController) Fetch(c *gin.Context) {

	NewsLetterSubscription, err := lc.NewsLetterSubscriptionUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewsLetterSubscription)
}
