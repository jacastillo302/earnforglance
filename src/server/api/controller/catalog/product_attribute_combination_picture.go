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

type ProductAttributeCombinationPictureController struct {
	ProductAttributeCombinationPictureUsecase domain.ProductAttributeCombinationPictureUsecase
	Env                                       *bootstrap.Env
}

func (tc *ProductAttributeCombinationPictureController) CreateMany(c *gin.Context) {
	var task []domain.ProductAttributeCombinationPicture
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

	err = tc.ProductAttributeCombinationPictureUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductAttributeCombinationPicture created successfully",
	})
}

func (tc *ProductAttributeCombinationPictureController) Create(c *gin.Context) {
	var task domain.ProductAttributeCombinationPicture
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

	err = tc.ProductAttributeCombinationPictureUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductAttributeCombinationPicture created successfully",
	})
}

func (tc *ProductAttributeCombinationPictureController) Update(c *gin.Context) {
	var task domain.ProductAttributeCombinationPicture
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

	err = tc.ProductAttributeCombinationPictureUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ProductAttributeCombinationPicture update successfully",
	})
}

func (tc *ProductAttributeCombinationPictureController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ProductAttributeCombinationPictureUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ProductAttributeCombinationPictureController) FetchByID(c *gin.Context) {
	ProductAttributeCombinationPictureID := c.Query("id")
	if ProductAttributeCombinationPictureID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ProductAttributeCombinationPicture, err := lc.ProductAttributeCombinationPictureUsecase.FetchByID(c, ProductAttributeCombinationPictureID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ProductAttributeCombinationPictureID})
		return
	}

	c.JSON(http.StatusOK, ProductAttributeCombinationPicture)
}

func (lc *ProductAttributeCombinationPictureController) Fetch(c *gin.Context) {

	ProductAttributeCombinationPicture, err := lc.ProductAttributeCombinationPictureUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProductAttributeCombinationPicture)
}
