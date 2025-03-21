package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/attributes"
	common "earnforglance/server/domain/common"

	"github.com/gin-gonic/gin"
)

type BaseAttributeController struct {
	BaseAttributeUsecase domain.BaseAttributeUsecase
	Env                  *bootstrap.Env
}

func (tc *BaseAttributeController) Create(c *gin.Context) {
	var task domain.BaseAttribute
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

	err = tc.BaseAttributeUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "BaseAttribute created successfully",
	})
}

func (tc *BaseAttributeController) Update(c *gin.Context) {
	var task domain.BaseAttribute
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

	err = tc.BaseAttributeUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "BaseAttribute update successfully",
	})
}

func (tc *BaseAttributeController) Delete(c *gin.Context) {
	var task domain.BaseAttribute
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

	err = tc.BaseAttributeUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "BaseAttribute update successfully",
	})
}

func (lc *BaseAttributeController) FetchByID(c *gin.Context) {
	BaseAttributeID := c.Query("id")
	if BaseAttributeID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	BaseAttribute, err := lc.BaseAttributeUsecase.FetchByID(c, BaseAttributeID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: BaseAttributeID})
		return
	}

	c.JSON(http.StatusOK, BaseAttribute)
}

func (lc *BaseAttributeController) Fetch(c *gin.Context) {

	BaseAttribute, err := lc.BaseAttributeUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, BaseAttribute)
}
