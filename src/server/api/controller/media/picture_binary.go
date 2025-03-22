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

type PictureBinaryController struct {
	PictureBinaryUsecase domain.PictureBinaryUsecase
	Env                  *bootstrap.Env
}

func (tc *PictureBinaryController) Create(c *gin.Context) {
	var task domain.PictureBinary
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

	err = tc.PictureBinaryUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PictureBinary created successfully",
	})
}

func (tc *PictureBinaryController) Update(c *gin.Context) {
	var task domain.PictureBinary
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

	err = tc.PictureBinaryUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "PictureBinary update successfully",
	})
}

func (tc *PictureBinaryController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.PictureBinaryUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *PictureBinaryController) FetchByID(c *gin.Context) {
	PictureBinaryID := c.Query("id")
	if PictureBinaryID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	PictureBinary, err := lc.PictureBinaryUsecase.FetchByID(c, PictureBinaryID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: PictureBinaryID})
		return
	}

	c.JSON(http.StatusOK, PictureBinary)
}

func (lc *PictureBinaryController) Fetch(c *gin.Context) {

	PictureBinary, err := lc.PictureBinaryUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, PictureBinary)
}
