package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	domain "earnforglance/server/domain/configuration"

	"github.com/gin-gonic/gin"
)

type SettingController struct {
	SettingUsecase domain.SettingUsecase
	Env            *bootstrap.Env
}

// NewSettingController creates a new instance of SettingController
func NewSettingController(settingUsecase domain.SettingUsecase, env *bootstrap.Env) *SettingController {
	return &SettingController{
		SettingUsecase: settingUsecase,
		Env:            env,
	}
}

func (tc *SettingController) CreateMany(c *gin.Context) {
	var task []domain.Setting
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

	err = tc.SettingUsecase.CreateMany(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Setting created successfully",
	})
}

func (tc *SettingController) Create(c *gin.Context) {
	var task domain.Setting
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

	err = tc.SettingUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Setting created successfully",
	})
}

func (tc *SettingController) Update(c *gin.Context) {
	var task domain.Setting
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

	err = tc.SettingUsecase.Update(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Setting update successfully",
	})
}

func (tc *SettingController) Delete(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	err := tc.SettingUsecase.Delete(c, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: ID})
		return
	}

	c.JSON(http.StatusOK, common.SuccessResponse{
		Message: "Record deleted successfully",
	})
}

func (lc *SettingController) FetchByID(c *gin.Context) {
	SettingID := c.Query("id")
	if SettingID == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "invalid ID format"})
		return
	}

	Setting, err := lc.SettingUsecase.FetchByID(c, SettingID)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: SettingID})
		return
	}

	c.JSON(http.StatusOK, Setting)
}

func (lc *SettingController) FetchByName(c *gin.Context) {
	SettingName := c.Query("name")
	if SettingName == "" {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "name is requiered"})
		return
	}

	Setting, err := lc.SettingUsecase.FetchByID(c, SettingName)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: SettingName})
		return
	}

	c.JSON(http.StatusOK, Setting)
}

func (lc *SettingController) Fetch(c *gin.Context) {

	Setting, err := lc.SettingUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Setting)
}

func (lc *SettingController) FetchByNames(c *gin.Context) {
	names := c.QueryArray("names") // Retrieve the "names" query parameter as an array
	if len(names) == 0 {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "names are required"})
		return
	}

	// Call the FetchByNames method from SettingUsecase
	settings, err := lc.SettingUsecase.FetchByNames(c, names)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, settings)
}
