package controller

import (
	"net/http"
	"time"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	response "earnforglance/server/domain/install"

	service "earnforglance/server/service/install"

	"github.com/gin-gonic/gin"
)

type InstallController struct {
	InstallUsecase response.InstallLogUsecase
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

	result, stores := service.InstallStores()
	if !result.Status {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(stores) == 0 {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: result.Details})
		return
	}

	err := lc.InstallUsecase.InstallStores(c, stores)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallCurrencies(c *gin.Context) {
	result, items := service.InstallCurrency()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No items to install"})
		return
	}

	err := lc.InstallUsecase.InstallCurrencies(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallMeasureDimension(c *gin.Context) {
	result, items := service.InstallMeasureDimension()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No items to install"})
		return
	}

	err := lc.InstallUsecase.InstallMeasureDimension(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallMeasureWeight(c *gin.Context) {
	result, items := service.InstallMeasureWeight()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No items to install"})
		return
	}

	err := lc.InstallUsecase.InstallMeasureWeight(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallTaxCategories(c *gin.Context) {
	result, items := service.InstallTaxCategory()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No items to install"})
		return
	}

	err := lc.InstallUsecase.InstallTaxCategory(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)

}

func (lc *InstallController) InstallLanguages(c *gin.Context) {
	result, items := service.InstallLanguages()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No items to install"})
		return
	}

	err := lc.InstallUsecase.InstallLanguages(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	/*
		defaultLanguage := service.GetDefaultLanguage()
		err := lc.InstallUsecase.InstallLanguages(c, defaultLanguage)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
			return
		}
	*/
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallStores(c *gin.Context) {

	result, stores := service.InstallStores()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(stores) == 0 || stores == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No stores to install"})
		return
	}

	err := lc.InstallUsecase.InstallStores(c, stores)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

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
