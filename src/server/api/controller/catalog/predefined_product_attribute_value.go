package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/catalog"
	common "earnforglance/server/domain/common"

	"github.com/gin-gonic/gin"
)

type PredefinedProductAttributeValueController struct {
	PredefinedProductAttributeValueUsecase domain.PredefinedProductAttributeValueUsecase
	Env                                    *bootstrap.Env
}

func (tc *PredefinedProductAttributeValueController) Create(c *gin.Context) {
	var task domain.PredefinedProductAttributeValue
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

	err = tc.PredefinedProductAttributeValueUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PredefinedProductAttributeValue created successfully",
	})
}

func (tc *PredefinedProductAttributeValueController) Update(c *gin.Context) {
	var task domain.PredefinedProductAttributeValue
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

	err = tc.PredefinedProductAttributeValueUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PredefinedProductAttributeValue update successfully",
	})
}

func (tc *PredefinedProductAttributeValueController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.PredefinedProductAttributeValueUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *PredefinedProductAttributeValueController) FetchByID(c *gin.Context) {
	PredefinedProductAttributeValueID := c.Query("id")
	if PredefinedProductAttributeValueID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	PredefinedProductAttributeValue, err := lc.PredefinedProductAttributeValueUsecase.FetchByID(c, PredefinedProductAttributeValueID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: PredefinedProductAttributeValueID})
		return
	}

	c.JSON(http.StatusOK, PredefinedProductAttributeValue)
}

func (lc *PredefinedProductAttributeValueController) Fetch(c *gin.Context) {

	PredefinedProductAttributeValue, err := lc.PredefinedProductAttributeValueUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, PredefinedProductAttributeValue)
}
