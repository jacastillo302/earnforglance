package controller

import (
	"net/http"
	"time"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	settings "earnforglance/server/domain/configuration"
	response "earnforglance/server/domain/install"
	stores "earnforglance/server/domain/stores"
	service "earnforglance/server/service/install"
	usecase "earnforglance/server/usecase/stores"

	"github.com/gin-gonic/gin"
)

type InstallController struct {
	InstallUsecase response.InstallLogUsecase
	SettingUsecase settings.SettingUsecase
	StoresUsecase  stores.StoreRepository
	Env            *bootstrap.Env
}

func (lc *InstallController) PingDatabase(c *gin.Context) {
	var result response.Install
	err := lc.InstallUsecase.PingDatabase(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	result.Status = true
	result.Details = "Database connection successful"
	result.CreatedOnUtc = time.Now()

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) FullInstall(c *gin.Context, sample bool) {

	result := service.InstallStores(lc.StoresUsecase)
	if !result.Status {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: result.Details})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallStores(c *gin.Context) {
	usecasevar := usecase.NewStoreUsecase(lc.StoresUsecase, time.Duration(0))
	result := service.InstallStores(usecasevar)
	if !result.Status {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: result.Details})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallMeasures(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallTaxCategories(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallLanguages(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallCurrencies(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallCountriesAndStates(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallShippingMethods(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallDeliveryDates(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallProductAvailabilityRanges(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallEmailAccounts(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallMessageTemplates(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallTopicTemplates(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallSettings(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallCustomersAndUsers(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallTopics(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallActivityLogTypes(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallProductTemplates(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallCategoryTemplates(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallManufacturerTemplates(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallScheduleTasks(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallReturnRequestReasons(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallReturnRequestActions(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallSampleData(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}
