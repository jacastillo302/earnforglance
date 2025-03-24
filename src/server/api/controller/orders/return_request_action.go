package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/orders"

	"github.com/gin-gonic/gin"
)

type ReturnRequestActionController struct {
	ReturnRequestActionUsecase domain.ReturnRequestActionUsecase
	Env                        *bootstrap.Env
}

func (tc *ReturnRequestActionController) CreateMany(c *gin.Context) {
	var task []domain.ReturnRequestAction
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

	err = tc.ReturnRequestActionUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ReturnRequestAction created successfully",
	})
}

func (tc *ReturnRequestActionController) Create(c *gin.Context) {
	var task domain.ReturnRequestAction
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

	err = tc.ReturnRequestActionUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ReturnRequestAction created successfully",
	})
}

func (tc *ReturnRequestActionController) Update(c *gin.Context) {
	var task domain.ReturnRequestAction
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

	err = tc.ReturnRequestActionUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ReturnRequestAction update successfully",
	})
}

func (tc *ReturnRequestActionController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ReturnRequestActionUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ReturnRequestActionController) FetchByID(c *gin.Context) {
	ReturnRequestActionID := c.Query("id")
	if ReturnRequestActionID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ReturnRequestAction, err := lc.ReturnRequestActionUsecase.FetchByID(c, ReturnRequestActionID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ReturnRequestActionID})
		return
	}

	c.JSON(http.StatusOK, ReturnRequestAction)
}

func (lc *ReturnRequestActionController) Fetch(c *gin.Context) {

	ReturnRequestAction, err := lc.ReturnRequestActionUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ReturnRequestAction)
}
