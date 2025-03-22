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

type ForumGroupController struct {
	ForumGroupUsecase domain.ForumGroupUsecase
	Env               *bootstrap.Env
}

func (tc *ForumGroupController) Create(c *gin.Context) {
	var task domain.ForumGroup
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

	err = tc.ForumGroupUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ForumGroup created successfully",
	})
}

func (tc *ForumGroupController) Update(c *gin.Context) {
	var task domain.ForumGroup
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

	err = tc.ForumGroupUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ForumGroup update successfully",
	})
}

func (tc *ForumGroupController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ForumGroupUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ForumGroupController) FetchByID(c *gin.Context) {
	ForumGroupID := c.Query("id")
	if ForumGroupID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ForumGroup, err := lc.ForumGroupUsecase.FetchByID(c, ForumGroupID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ForumGroupID})
		return
	}

	c.JSON(http.StatusOK, ForumGroup)
}

func (lc *ForumGroupController) Fetch(c *gin.Context) {

	ForumGroup, err := lc.ForumGroupUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ForumGroup)
}
