package controller

import (
	"net/http"
	"time"

	"earnforglance/server/bootstrap"
	common "earnforglance/server/domain/common"
	response "earnforglance/server/domain/install"
	service "earnforglance/server/service/install"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No PermissionRecord to install"})
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
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No Currencies to install"})
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
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No MeasureDimension to install"})
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
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No MeasureWeight to install"})
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
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No TaxCategories to install"})
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
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No Languages to install"})
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
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No Settings to install"})
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
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No DeliveryDate to install"})
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
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No ProductAvailabilityRange to install"})
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

func (lc *InstallController) InstallMessageTemplate(c *gin.Context) {
	result, items := service.InstallMessageTemplate()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No MessageTemplate to install"})
		return
	}

	err := lc.InstallUsecase.InstallMessageTemplate(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallTopicTemplate(c *gin.Context) {
	result, items := service.InstallTopicTemplate()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No TopicTemplate to install"})
		return
	}

	err := lc.InstallUsecase.InstallTopicTemplate(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	result, topics := service.InstallTopic()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	for index, _ := range topics {
		topics[index].TopicTemplateID = items[0].ID

	}

	if len(topics) == 0 || topics == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No Topics to install"})
		return
	}

	err = lc.InstallUsecase.InstallTopic(c, topics)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallCustomerRole(c *gin.Context) {

	result, roles := service.InstallCustomerRole()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(roles) == 0 || roles == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No CustomerRole to install"})
		return
	}

	err := lc.InstallUsecase.InstallCustomerRole(c, roles)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	result, customer := service.InstallCustomer()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(customer) == 0 || customer == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No Customer to install"})
		return
	}

	customer[0].CreatedOnUtc = time.Now()
	customer[0].LastActivityDateUtc = time.Now()
	customer[0].LastIpAddress = c.ClientIP()
	customer[0].LastLoginDateUtc = nil
	customer[0].CannotLoginUntilDateUtc = nil

	err = lc.InstallUsecase.InstallCustomer(c, customer)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	result, custroles := service.InstallCustomerCustomerRoleMapping(customer[0].ID, roles)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	err = lc.InstallUsecase.InstallCustomerCustomerRoleMapping(c, custroles)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	result, cusp := service.InstallCustomerPassword(customer[0].ID, customer[0].Email)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(cusp) == 0 || cusp == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No CustomerPassword to install"})
		return
	}

	err = lc.InstallUsecase.InstallCustomerPassword(c, cusp)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	///Adresses
	result, adress := service.InstallCustomerAddress()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(adress) == 0 || adress == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No CustomerAddress to install"})
		return
	}

	adress[0].CreatedOnUtc = time.Now()
	adress[0].Email = customer[0].Email
	adress[0].CountryID = &customer[0].CountryID
	adress[0].StateProvinceID = &customer[0].StateProvinceID

	err = lc.InstallUsecase.InstallCustomerAddress(c, adress)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	billingAddressID, err := primitive.ObjectIDFromHex(adress[0].ID.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Invalid address ID: " + err.Error()})
		return
	}
	customer[0].BillingAddressID = &billingAddressID

	shippingAddressID, err := primitive.ObjectIDFromHex(adress[0].ID.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Invalid address ID: " + err.Error()})
		return
	}
	customer[0].ShippingAddressID = &shippingAddressID

	err = lc.InstallUsecase.UpdateCustomer(c, customer[0])
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	result, adressmap := service.InstallCustomerAddressMapping(customer[0].ID, adress[0].ID)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	err = lc.InstallUsecase.InstallCustomerAddressMapping(c, adressmap)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallActivityLogType(c *gin.Context) {

	result, items := service.InstallActivityLogType()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No MessageTemplate to install"})
		return
	}

	err := lc.InstallUsecase.InstallActivityLogType(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallProductTemplate(c *gin.Context) {
	result, items := service.InstallProductTemplate()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No MessageTemplate to install"})
		return
	}

	err := lc.InstallUsecase.InstallProductTemplate(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallCategoryTemplate(c *gin.Context) {
	result, items := service.InstallCategoryTemplate()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No MessageTemplate to install"})
		return
	}

	err := lc.InstallUsecase.InstallCategoryTemplate(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallManufacturerTemplate(c *gin.Context) {
	result, items := service.InstallManufacturerTemplate()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No MessageTemplate to install"})
		return
	}

	err := lc.InstallUsecase.InstallManufacturerTemplate(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallScheduleTask(c *gin.Context) {
	result, items := service.InstallScheduleTask()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No ScheduleTask to install"})
		return
	}

	err := lc.InstallUsecase.InstallScheduleTask(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallReturnRequestReason(c *gin.Context) {
	result, items := service.InstallReturnRequestReason()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No ReturnRequestReason to install"})
		return
	}

	err := lc.InstallUsecase.InstallReturnRequestReason(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallReturnRequestAction(c *gin.Context) {
	result, items := service.InstallReturnRequestAction()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if len(items) == 0 || items == nil {
		c.JSON(http.StatusNoContent, common.ErrorResponse{Message: "No ReturnRequestAction to install"})
		return
	}

	err := lc.InstallUsecase.InstallReturnRequestAction(c, items)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallSampleData(c *gin.Context) {
	var result response.Install
	c.JSON(http.StatusOK, result)
}
