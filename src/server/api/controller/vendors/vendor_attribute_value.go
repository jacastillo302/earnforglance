package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/vendors"

	"github.com/gin-gonic/gin"
)

type VendorAttributeValueController struct {
	VendorAttributeValueUsecase domain.VendorAttributeValueUsecase
	Env                         *bootstrap.Env
}

func (tc *VendorAttributeValueController) Create(c *gin.Context) {
	var task domain.VendorAttributeValue
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

	err = tc.VendorAttributeValueUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "VendorAttributeValue created successfully",
	})
}

func (tc *VendorAttributeValueController) Update(c *gin.Context) {
	var task domain.VendorAttributeValue
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

	err = tc.VendorAttributeValueUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "VendorAttributeValue update successfully",
	})
}

func (tc *VendorAttributeValueController) Delete(c *gin.Context) {
	var task domain.VendorAttributeValue
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

	err = tc.VendorAttributeValueUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "VendorAttributeValue update successfully",
	})
}

func (lc *VendorAttributeValueController) FetchByID(c *gin.Context) {
	VendorAttributeValueID := c.Query("id")
	if VendorAttributeValueID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	VendorAttributeValue, err := lc.VendorAttributeValueUsecase.FetchByID(c, VendorAttributeValueID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: VendorAttributeValueID})
		return
	}

	c.JSON(http.StatusOK, VendorAttributeValue)
}

func (lc *VendorAttributeValueController) Fetch(c *gin.Context) {

	VendorAttributeValue, err := lc.VendorAttributeValueUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, VendorAttributeValue)
}
