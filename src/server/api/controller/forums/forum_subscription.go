package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/forums"

	"github.com/gin-gonic/gin"
)

type ForumSubscriptionController struct {
	ForumSubscriptionUsecase domain.ForumSubscriptionUsecase
	Env                      *bootstrap.Env
}

func (tc *ForumSubscriptionController) Create(c *gin.Context) {
	var task domain.ForumSubscription
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

	err = tc.ForumSubscriptionUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ForumSubscription created successfully",
	})
}

func (tc *ForumSubscriptionController) Update(c *gin.Context) {
	var task domain.ForumSubscription
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

	err = tc.ForumSubscriptionUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ForumSubscription update successfully",
	})
}

func (tc *ForumSubscriptionController) Delete(c *gin.Context) {
	var task domain.ForumSubscription
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

	err = tc.ForumSubscriptionUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ForumSubscription update successfully",
	})
}

func (lc *ForumSubscriptionController) FetchByID(c *gin.Context) {
	ForumSubscriptionID := c.Query("id")
	if ForumSubscriptionID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ForumSubscription, err := lc.ForumSubscriptionUsecase.FetchByID(c, ForumSubscriptionID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ForumSubscriptionID})
		return
	}

	c.JSON(http.StatusOK, ForumSubscription)
}

func (lc *ForumSubscriptionController) Fetch(c *gin.Context) {

	ForumSubscription, err := lc.ForumSubscriptionUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ForumSubscription)
}
