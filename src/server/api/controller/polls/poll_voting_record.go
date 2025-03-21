package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/polls"

	"github.com/gin-gonic/gin"
)

type PollVotingRecordController struct {
	PollVotingRecordUsecase domain.PollVotingRecordUsecase
	Env                     *bootstrap.Env
}

func (tc *PollVotingRecordController) Create(c *gin.Context) {
	var task domain.PollVotingRecord
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

	err = tc.PollVotingRecordUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PollVotingRecord created successfully",
	})
}

func (tc *PollVotingRecordController) Update(c *gin.Context) {
	var task domain.PollVotingRecord
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

	err = tc.PollVotingRecordUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PollVotingRecord update successfully",
	})
}

func (tc *PollVotingRecordController) Delete(c *gin.Context) {
	var task domain.PollVotingRecord
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

	err = tc.PollVotingRecordUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PollVotingRecord update successfully",
	})
}

func (lc *PollVotingRecordController) FetchByID(c *gin.Context) {
	PollVotingRecordID := c.Query("id")
	if PollVotingRecordID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	PollVotingRecord, err := lc.PollVotingRecordUsecase.FetchByID(c, PollVotingRecordID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: PollVotingRecordID})
		return
	}

	c.JSON(http.StatusOK, PollVotingRecord)
}

func (lc *PollVotingRecordController) Fetch(c *gin.Context) {

	PollVotingRecord, err := lc.PollVotingRecordUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, PollVotingRecord)
}
