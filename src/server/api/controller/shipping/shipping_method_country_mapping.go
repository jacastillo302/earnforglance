package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/shipping"

	"github.com/gin-gonic/gin"
)

type ShippingMethodCountryMappingController struct {
	ShippingMethodCountryMappingUsecase domain.ShippingMethodCountryMappingUsecase
	Env                                 *bootstrap.Env
}

func (tc *ShippingMethodCountryMappingController) CreateMany(c *gin.Context) {
	var task []domain.ShippingMethodCountryMapping
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

	err = tc.ShippingMethodCountryMappingUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ShippingMethodCountryMapping created successfully",
	})
}

func (tc *ShippingMethodCountryMappingController) Create(c *gin.Context) {
	var task domain.ShippingMethodCountryMapping
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

	err = tc.ShippingMethodCountryMappingUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ShippingMethodCountryMapping created successfully",
	})
}

func (tc *ShippingMethodCountryMappingController) Update(c *gin.Context) {
	var task domain.ShippingMethodCountryMapping
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

	err = tc.ShippingMethodCountryMappingUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "ShippingMethodCountryMapping update successfully",
	})
}

func (tc *ShippingMethodCountryMappingController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.ShippingMethodCountryMappingUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *ShippingMethodCountryMappingController) FetchByID(c *gin.Context) {
	ShippingMethodCountryMappingID := c.Query("id")
	if ShippingMethodCountryMappingID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	ShippingMethodCountryMapping, err := lc.ShippingMethodCountryMappingUsecase.FetchByID(c, ShippingMethodCountryMappingID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ShippingMethodCountryMappingID})
		return
	}

	c.JSON(http.StatusOK, ShippingMethodCountryMapping)
}

func (lc *ShippingMethodCountryMappingController) Fetch(c *gin.Context) {

	ShippingMethodCountryMapping, err := lc.ShippingMethodCountryMappingUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ShippingMethodCountryMapping)
}
