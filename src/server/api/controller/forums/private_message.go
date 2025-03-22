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

type PrivateMessageController struct {
	PrivateMessageUsecase domain.PrivateMessageUsecase
	Env                   *bootstrap.Env
}

func (tc *PrivateMessageController) Create(c *gin.Context) {
	var task domain.PrivateMessage
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

	err = tc.PrivateMessageUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PrivateMessage created successfully",
	})
}

func (tc *PrivateMessageController) Update(c *gin.Context) {
	var task domain.PrivateMessage
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

	err = tc.PrivateMessageUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PrivateMessage update successfully",
	})
}

func (tc *PrivateMessageController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.PrivateMessageUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *PrivateMessageController) FetchByID(c *gin.Context) {
	PrivateMessageID := c.Query("id")
	if PrivateMessageID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	PrivateMessage, err := lc.PrivateMessageUsecase.FetchByID(c, PrivateMessageID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: PrivateMessageID})
		return
	}

	c.JSON(http.StatusOK, PrivateMessage)
}

func (lc *PrivateMessageController) Fetch(c *gin.Context) {

	PrivateMessage, err := lc.PrivateMessageUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, PrivateMessage)
}
