package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/polls" // changed [dirname] to polls

	"github.com/gin-gonic/gin"
)

type PollController struct { // changed Poll to Poll
	PollUsecase domain.PollUsecase // changed PollUsecase to PollUsecase
	Env         *bootstrap.Env
}

func (tc *PollController) CreateMany(c *gin.Context) {
	var task []domain.Poll
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

	err = tc.PollUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Poll created successfully",
	})
}

func (tc *PollController) Create(c *gin.Context) {
	var task domain.Poll // changed Poll to Poll
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
	err = tc.PollUsecase.Create(c, &task) // changed PollUsecase to PollUsecase
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Poll created successfully", // changed Poll to Poll
	})
}

func (tc *PollController) Update(c *gin.Context) {
	var task domain.Poll // changed Poll to Poll
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
	err = tc.PollUsecase.Update(c, &task) // changed PollUsecase to PollUsecase
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Poll update successfully", // changed Poll to Poll
	})
}

func (tc *PollController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.PollUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *PollController) FetchByID(c *gin.Context) {
	pollID := c.Query("id") // changed PollID to pollID
	if pollID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}
	poll, err := lc.PollUsecase.FetchByID(c, pollID) // changed PollUsecase to PollUsecase
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: pollID})
		return
	}
	c.JSON(http.StatusOK, poll) // changed Poll to poll
}

func (lc *PollController) Fetch(c *gin.Context) {
	poll, err := lc.PollUsecase.Fetch(c) // changed Poll to Poll
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, poll) // changed Poll to poll
}
