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

type ForumTopicController struct {
	ForumTopicUsecase domain.ForumTopicUsecase
	Env               *bootstrap.Env
}

func (tc *ForumTopicController) Create(c *gin.Context) {
	var task domain.ForumTopic
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

	err = tc.ForumTopicUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ForumTopic created successfully",
	})
}

func (tc *ForumTopicController) Update(c *gin.Context) {
	var task domain.ForumTopic
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

	err = tc.ForumTopicUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ForumTopic update successfully",
	})
}

func (tc *ForumTopicController) Delete(c *gin.Context) {
	var task domain.ForumTopic
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

	err = tc.ForumTopicUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ForumTopic update successfully",
	})
}

func (lc *ForumTopicController) FetchByID(c *gin.Context) {
	ForumTopicID := c.Query("id")
	if ForumTopicID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ForumTopic, err := lc.ForumTopicUsecase.FetchByID(c, ForumTopicID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ForumTopicID})
		return
	}

	c.JSON(http.StatusOK, ForumTopic)
}

func (lc *ForumTopicController) Fetch(c *gin.Context) {

	ForumTopic, err := lc.ForumTopicUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ForumTopic)
}
