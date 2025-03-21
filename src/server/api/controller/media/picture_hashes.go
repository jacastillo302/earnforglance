package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/media"

	"github.com/gin-gonic/gin"
)

type PictureHashesController struct {
	PictureHashesUsecase domain.PictureHashesUsecase
	Env                  *bootstrap.Env
}

func (tc *PictureHashesController) Create(c *gin.Context) {
	var task domain.PictureHashes
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

	err = tc.PictureHashesUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PictureHashes created successfully",
	})
}

func (tc *PictureHashesController) Update(c *gin.Context) {
	var task domain.PictureHashes
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

	err = tc.PictureHashesUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PictureHashes update successfully",
	})
}

func (tc *PictureHashesController) Delete(c *gin.Context) {
	var task domain.PictureHashes
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

	err = tc.PictureHashesUsecase.Delete(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PictureHashes update successfully",
	})
}

func (lc *PictureHashesController) FetchByID(c *gin.Context) {
	PictureHashesID := c.Query("id")
	if PictureHashesID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	PictureHashes, err := lc.PictureHashesUsecase.FetchByID(c, PictureHashesID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: PictureHashesID})
		return
	}

	c.JSON(http.StatusOK, PictureHashes)
}

func (lc *PictureHashesController) Fetch(c *gin.Context) {

	PictureHashes, err := lc.PictureHashesUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, PictureHashes)
}
