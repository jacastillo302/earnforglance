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

type VendorNoteController struct {
	VendorNoteUsecase domain.VendorNoteUsecase
	Env               *bootstrap.Env
}

func (tc *VendorNoteController) Create(c *gin.Context) {
	var task domain.VendorNote
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

	err = tc.VendorNoteUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "VendorNote created successfully",
	})
}

func (tc *VendorNoteController) Update(c *gin.Context) {
	var task domain.VendorNote
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

	err = tc.VendorNoteUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "VendorNote update successfully",
	})
}

func (tc *VendorNoteController) Delete(c *gin.Context) {
	var task domain.VendorNote
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

	err = tc.VendorNoteUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "VendorNote update successfully",
	})
}

func (lc *VendorNoteController) FetchByID(c *gin.Context) {
	VendorNoteID := c.Query("id")
	if VendorNoteID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	VendorNote, err := lc.VendorNoteUsecase.FetchByID(c, VendorNoteID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: VendorNoteID})
		return
	}

	c.JSON(http.StatusOK, VendorNote)
}

func (lc *VendorNoteController) Fetch(c *gin.Context) {

	VendorNote, err := lc.VendorNoteUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, VendorNote)
}
