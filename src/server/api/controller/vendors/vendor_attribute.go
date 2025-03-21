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

type VendorAttributeController struct {
	VendorAttributeUsecase domain.VendorAttributeUsecase
	Env                    *bootstrap.Env
}

func (tc *VendorAttributeController) Create(c *gin.Context) {
	var task domain.VendorAttribute
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

	err = tc.VendorAttributeUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "VendorAttribute created successfully",
	})
}

func (tc *VendorAttributeController) Update(c *gin.Context) {
	var task domain.VendorAttribute
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

	err = tc.VendorAttributeUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "VendorAttribute update successfully",
	})
}

func (tc *VendorAttributeController) Delete(c *gin.Context) {
	var task domain.VendorAttribute
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

	err = tc.VendorAttributeUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "VendorAttribute update successfully",
	})
}

func (lc *VendorAttributeController) FetchByID(c *gin.Context) {
	VendorAttributeID := c.Query("id")
	if VendorAttributeID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	VendorAttribute, err := lc.VendorAttributeUsecase.FetchByID(c, VendorAttributeID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: VendorAttributeID})
		return
	}

	c.JSON(http.StatusOK, VendorAttribute)
}

func (lc *VendorAttributeController) Fetch(c *gin.Context) {

	VendorAttribute, err := lc.VendorAttributeUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, VendorAttribute)
}
