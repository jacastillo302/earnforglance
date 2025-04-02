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

func (lc *InstallController) InstallPermissionRecord(c *gin.Context) {
	result, items := service.InstallPermissionRecord()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No items to install"})
		return
	}

	err := lc.InstallUsecase.InstallPermissionRecord(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
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

	for _, item := range items {
		result, rows := service.InstallLocaleStringResource(item.ID.Hex(), item.LanguageCulture)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}

		err := lc.InstallUsecase.InstallLocaleStringResource(c, rows)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
			return
		}

	}

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

func (lc *InstallController) InstallSettings(c *gin.Context) {
	result, items := service.InstallSettings()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No stores to install"})
		return
	}

	err := lc.InstallUsecase.InstallSettings(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallCountries(c *gin.Context) {
	result, items := service.InstallCountries()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No countries to install"})
		return
	}

	err := lc.InstallUsecase.InstallCountries(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	result, itemsb := service.InstallStateProvince()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	err = lc.InstallUsecase.InstallStateProvince(c, itemsb)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallShippingMethod(c *gin.Context) {
	result, items := service.InstallShippingMethod()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No ShippingMethod to install"})
		return
	}

	err := lc.InstallUsecase.InstallShippingMethod(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallDeliveryDate(c *gin.Context) {
	result, items := service.InstallDeliveryDate()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No ShippingMethod to install"})
		return
	}

	err := lc.InstallUsecase.InstallDeliveryDate(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallProductAvailabilityRange(c *gin.Context) {
	result, items := service.InstallProductAvailabilityRange()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No ShippingMethod to install"})
		return
	}

	err := lc.InstallUsecase.InstallProductAvailabilityRange(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallEmailAccount(c *gin.Context) {
	result, items := service.InstallEmailAccount()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No EmailAccount to install"})
		return
	}

	err := lc.InstallUsecase.InstallEmailAccount(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}
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
