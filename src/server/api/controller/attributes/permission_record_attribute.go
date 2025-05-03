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

type PermisionRecordAttributeController struct {
	PermisionRecordAttributeUsecase domain.PermisionRecordAttributeUsecase
	Env                             *bootstrap.Env
}

func (tc *PermisionRecordAttributeController) CreateMany(c *gin.Context) {
	var task []domain.PermisionRecordAttribute
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

	err = tc.PermisionRecordAttributeUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PermisionRecordAttribute created successfully",
	})
}

func (tc *PermisionRecordAttributeController) Create(c *gin.Context) {
	var task domain.PermisionRecordAttribute
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

	err = tc.PermisionRecordAttributeUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PermisionRecordAttribute created successfully",
	})
}

func (tc *PermisionRecordAttributeController) Update(c *gin.Context) {
	var task domain.PermisionRecordAttribute
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

	err = tc.PermisionRecordAttributeUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PermisionRecordAttribute update successfully",
	})
}

func (tc *PermisionRecordAttributeController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.PermisionRecordAttributeUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *PermisionRecordAttributeController) FetchByID(c *gin.Context) {
	PermisionRecordAttributeID := c.Query("id")
	if PermisionRecordAttributeID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	PermisionRecordAttribute, err := lc.PermisionRecordAttributeUsecase.FetchByID(c, PermisionRecordAttributeID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: PermisionRecordAttributeID})
		return
	}

	c.JSON(http.StatusOK, PermisionRecordAttribute)
}

func (lc *PermisionRecordAttributeController) Fetch(c *gin.Context) {

	PermisionRecordAttribute, err := lc.PermisionRecordAttributeUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, PermisionRecordAttribute)
}
