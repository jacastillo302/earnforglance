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

type PermisionRecordAttributeValueController struct {
	PermisionRecordAttributeValueUsecase domain.PermisionRecordAttributeValueUsecase
	Env                                  *bootstrap.Env
}

func (tc *PermisionRecordAttributeValueController) CreateMany(c *gin.Context) {
	var task []domain.PermisionRecordAttributeValue
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

	err = tc.PermisionRecordAttributeValueUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PermisionRecordAttributeValue created successfully",
	})
}

func (tc *PermisionRecordAttributeValueController) Create(c *gin.Context) {
	var task domain.PermisionRecordAttributeValue
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

	err = tc.PermisionRecordAttributeValueUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PermisionRecordAttributeValue created successfully",
	})
}

func (tc *PermisionRecordAttributeValueController) Update(c *gin.Context) {
	var task domain.PermisionRecordAttributeValue
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

	err = tc.PermisionRecordAttributeValueUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PermisionRecordAttributeValue update successfully",
	})
}

func (tc *PermisionRecordAttributeValueController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.PermisionRecordAttributeValueUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *PermisionRecordAttributeValueController) FetchByID(c *gin.Context) {
	PermisionRecordAttributeValueID := c.Query("id")
	if PermisionRecordAttributeValueID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	PermisionRecordAttributeValue, err := lc.PermisionRecordAttributeValueUsecase.FetchByID(c, PermisionRecordAttributeValueID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: PermisionRecordAttributeValueID})
		return
	}

	c.JSON(http.StatusOK, PermisionRecordAttributeValue)
}

func (lc *PermisionRecordAttributeValueController) Fetch(c *gin.Context) {

	PermisionRecordAttributeValue, err := lc.PermisionRecordAttributeValueUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, PermisionRecordAttributeValue)
}
