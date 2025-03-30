package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/customers"

	"github.com/gin-gonic/gin"
)

type ExternalAuthenticationRecordController struct {
	ExternalAuthenticationRecordUsecase domain.ExternalAuthenticationRecordUsecase
	Env                                 *bootstrap.Env
}

func (tc *ExternalAuthenticationRecordController) CreateMany(c *gin.Context) {
	var task []domain.ExternalAuthenticationRecord
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

	err = tc.ExternalAuthenticationRecordUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ExternalAuthenticationRecord created successfully",
	})
}

func (tc *ExternalAuthenticationRecordController) Create(c *gin.Context) {
	var task domain.ExternalAuthenticationRecord
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

	err = tc.ExternalAuthenticationRecordUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ExternalAuthenticationRecord created successfully",
	})
}

func (tc *ExternalAuthenticationRecordController) Update(c *gin.Context) {
	var task domain.ExternalAuthenticationRecord
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

	err = tc.ExternalAuthenticationRecordUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ExternalAuthenticationRecord updated successfully",
	})
}

func (tc *ExternalAuthenticationRecordController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ExternalAuthenticationRecordUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ExternalAuthenticationRecordController) FetchByID(c *gin.Context) {
	ExternalAuthenticationRecordID := c.Query("id")
	if ExternalAuthenticationRecordID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ExternalAuthenticationRecord, err := lc.ExternalAuthenticationRecordUsecase.FetchByID(c, ExternalAuthenticationRecordID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ExternalAuthenticationRecordID})
		return
	}

	c.JSON(http.StatusOK, ExternalAuthenticationRecord)
}

func (lc *ExternalAuthenticationRecordController) Fetch(c *gin.Context) {
	ExternalAuthenticationRecord, err := lc.ExternalAuthenticationRecordUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ExternalAuthenticationRecord)
}
