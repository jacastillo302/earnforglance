package controller

import (
	"encoding/json"
	"io"
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

func (lc *InstallController) FullInstall(c *gin.Context) {

	var result response.Install
	var model response.InstallModel

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Failed to read request body"})
		return
	}

	err = json.Unmarshal(body, &model)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Invalid request body"})
		return
	}

	err = lc.InstallUsecase.PingDatabase(c)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorResponse{Message: err.Error()})
		return
	}

	result = PermissionRecord(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = Currencies(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = MeasureDimension(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = TaxCategories(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = Languages(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = Stores(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = Settings(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = Countries(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = ShippingMethod(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = DeliveryDate(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = ProductAvailabilityRange(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = EmailAccount(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = MessageTemplate(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = TopicTemplate(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = CustomerRole(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = ProductTemplate(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = CategoryTemplate(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = ManufacturerTemplate(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = ScheduleTask(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = ReturnRequestReason(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = ReturnRequestAction(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	result = ActivityLogType(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	result = CustomerAdmin(c, lc, model)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}

	if model.SampleData {
		result = CustomerData(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = CheckoutAttribute(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = SpecificationAttribute(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = ProductAttribute(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = Picture(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = Category(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = Manufacturer(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = Warehouse(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = Vendor(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = Affiliate(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = Forum(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = Discount(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = BlogPost(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = Poll(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = NewsItem(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = SearchTerm(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = Download(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = GdprConsent(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = Product(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
		result = Order(c, lc, model.SampleData)
		if !result.Status {
			c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
			return
		}
	}

	result.Status = true
	result.Details = "Full Install Completed, login with user name: " + model.AdminEmail

	c.JSON(http.StatusOK, result)
}

func (lc *InstallController) InstallPermissionRecord(c *gin.Context) {
	result := PermissionRecord(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func PermissionRecord(c *gin.Context, lc *InstallController) response.Install {

	result, items := service.InstallPermissionRecord()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No PermissionRecord to install"
		return result
	}

	err := lc.InstallUsecase.InstallPermissionRecord(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallCurrencies(c *gin.Context) {
	result := Currencies(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Currencies(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallCurrency()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No Currencies to install"
		return result
	}

	err := lc.InstallUsecase.InstallCurrencies(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallMeasureDimension(c *gin.Context) {
	result := MeasureDimension(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func MeasureDimension(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallMeasureDimension()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No MeasureDimension to install"
		return result
	}

	err := lc.InstallUsecase.InstallMeasureDimension(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, weights := service.InstallMeasureWeight()
	if !result.Status {
		return result
	}

	if len(weights) == 0 || weights == nil {
		result.Status = false
		result.Details = "No MeasureWeight to install"
		return result
	}

	err = lc.InstallUsecase.InstallMeasureWeight(c, weights)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallTaxCategories(c *gin.Context) {
	result := TaxCategories(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)

}

func TaxCategories(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallTaxCategory()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No TaxCategories to install"
		return result
	}

	err := lc.InstallUsecase.InstallTaxCategory(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallLanguages(c *gin.Context) {
	result := Languages(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Languages(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallLanguages()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No Languages to install"
		return result
	}

	err := lc.InstallUsecase.InstallLanguages(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	for _, item := range items {
		result, rows := service.InstallLocaleStringResource(item.ID.Hex(), item.LanguageCulture)
		if !result.Status {
			return result
		}

		err := lc.InstallUsecase.InstallLocaleStringResource(c, rows)
		if err != nil {
			result.Status = false
			result.Details = err.Error()
			return result
		}

	}

	return result
}

func (lc *InstallController) InstallStores(c *gin.Context) {
	result := Stores(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Stores(c *gin.Context, lc *InstallController) response.Install {

	result, stores := service.InstallStores()
	if !result.Status {
		return result
	}

	if len(stores) == 0 || stores == nil {
		result.Status = false
		result.Details = "No Stores to install"
		return result
	}

	err := lc.InstallUsecase.InstallStores(c, stores)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallSettings(c *gin.Context) {
	result := Settings(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Settings(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallSettings()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No Settings to install"
		return result
	}

	err := lc.InstallUsecase.InstallSettings(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallCountries(c *gin.Context) {
	result := Countries(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Countries(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallCountries()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No Countries to install"
		return result
	}

	err := lc.InstallUsecase.InstallCountries(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, itemsb := service.InstallStateProvince()
	if !result.Status {
		return result
	}

	err = lc.InstallUsecase.InstallStateProvince(c, itemsb)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallShippingMethod(c *gin.Context) {
	result := ShippingMethod(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func ShippingMethod(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallShippingMethod()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No ShippingMethod to install"
		return result
	}

	err := lc.InstallUsecase.InstallShippingMethod(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallDeliveryDate(c *gin.Context) {
	result := DeliveryDate(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeliveryDate(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallDeliveryDate()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No DeliveryDate to install"
		return result
	}

	err := lc.InstallUsecase.InstallDeliveryDate(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallProductAvailabilityRange(c *gin.Context) {
	result := ProductAvailabilityRange(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func ProductAvailabilityRange(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallProductAvailabilityRange()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No ProductAvailabilityRange to install"
		return result
	}

	err := lc.InstallUsecase.InstallProductAvailabilityRange(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallEmailAccount(c *gin.Context) {
	result := EmailAccount(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func EmailAccount(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallEmailAccount()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No EmailAccount to install"
		return result
	}

	err := lc.InstallUsecase.InstallEmailAccount(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}
	return result
}

func (lc *InstallController) InstallMessageTemplate(c *gin.Context) {
	result := MessageTemplate(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func MessageTemplate(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallMessageTemplate()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No MessageTemplate to install"
		return result
	}

	err := lc.InstallUsecase.InstallMessageTemplate(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}
	return result
}

func (lc *InstallController) InstallTopicTemplate(c *gin.Context) {
	result := TopicTemplate(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func TopicTemplate(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallTopicTemplate()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No TopicTemplate to install"
		return result
	}

	err := lc.InstallUsecase.InstallTopicTemplate(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, topics := service.InstallTopic()
	if !result.Status {
		return result
	}

	for index := range topics {
		topics[index].TopicTemplateID = items[0].ID

	}

	if len(topics) == 0 || topics == nil {
		result.Status = false
		result.Details = "No topics to install"
		return result
	}

	err = lc.InstallUsecase.InstallTopic(c, topics)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallCustomerRole(c *gin.Context) {

	result := CustomerRole(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CustomerRole(c *gin.Context, lc *InstallController) response.Install {

	result, roles := service.InstallCustomerRole()
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return result
	}

	if len(roles) == 0 || roles == nil {
		result.Status = false
		result.Details = "No CustomerRole to install"
		return result
	}

	err := lc.InstallUsecase.InstallCustomerRole(c, roles)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}
	return result
}

func (lc *InstallController) InstallActivityLogType(c *gin.Context) {

	result := ActivityLogType(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func ActivityLogType(c *gin.Context, lc *InstallController) response.Install {

	result, items := service.InstallActivityLogType()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No ActivityLogType to install"
		return result
	}

	err := lc.InstallUsecase.InstallActivityLogType(c, items)
	if !result.Status {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, acty := service.InstallActivityLog()
	if !result.Status {
		return result
	}

	if len(acty) == 0 || acty == nil {
		result.Status = false
		result.Details = "No ActivityLog to install"
		return result
	}

	err = lc.InstallUsecase.InstallActivityLog(c, acty)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallProductTemplate(c *gin.Context) {
	result := ProductTemplate(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func ProductTemplate(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallProductTemplate()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No ProductTemplate to install"
		return result
	}

	err := lc.InstallUsecase.InstallProductTemplate(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallCategoryTemplate(c *gin.Context) {
	result := CategoryTemplate(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CategoryTemplate(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallCategoryTemplate()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No CategoryTemplate to install"
		return result
	}

	err := lc.InstallUsecase.InstallCategoryTemplate(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallManufacturerTemplate(c *gin.Context) {
	result := ManufacturerTemplate(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func ManufacturerTemplate(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallManufacturerTemplate()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No ManufacturerTemplate to install"
		return result
	}

	err := lc.InstallUsecase.InstallManufacturerTemplate(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallScheduleTask(c *gin.Context) {
	result := ScheduleTask(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func ScheduleTask(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallScheduleTask()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No ScheduleTask to install"
		return result
	}

	err := lc.InstallUsecase.InstallScheduleTask(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallReturnRequestReason(c *gin.Context) {
	result := ReturnRequestReason(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func ReturnRequestReason(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallReturnRequestReason()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No ReturnRequestReason to install"
		return result
	}

	err := lc.InstallUsecase.InstallReturnRequestReason(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallReturnRequestAction(c *gin.Context) {
	result := ReturnRequestAction(c, lc)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func ReturnRequestAction(c *gin.Context, lc *InstallController) response.Install {
	result, items := service.InstallReturnRequestAction()
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No ReturnRequestAction to install"
		return result
	}

	err := lc.InstallUsecase.InstallReturnRequestAction(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallCustomerSampleData(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := CustomerData(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CustomerAdmin(c *gin.Context, lc *InstallController, model response.InstallModel) response.Install {
	var result response.Install

	result, customer := service.InstallCustomer(false)
	if !result.Status {
		return result
	}

	if len(customer) == 0 || customer == nil {
		result.Status = false
		result.Details = "No customer to install"
		return result
	}

	customer[0].Username = model.AdminEmail
	customer[0].Email = model.AdminEmail

	err := lc.InstallUsecase.InstallCustomer(c, customer)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, roles := service.InstallCustomerRole()
	if !result.Status {
		return result
	}

	result, rolesm := service.InstallCustomerCustomerRoleMapping(customer[0].ID, roles)
	if !result.Status {
		return result
	}

	err = lc.InstallUsecase.InstallCustomerCustomerRoleMapping(c, rolesm)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	///Adresses
	result, adress := service.InstallAddress(false)
	if !result.Status {
		return result
	}

	if len(adress) == 0 || adress == nil {
		result.Status = false
		result.Details = "No Adresses to install"
		return result
	}

	err = lc.InstallUsecase.InstallAddress(c, adress)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	for i := 0; i < len(customer); i++ {

		customer[i].CreatedOnUtc = time.Now()
		customer[i].LastActivityDateUtc = time.Now()
		customer[i].LastIpAddress = c.ClientIP()
		customer[i].LastLoginDateUtc = nil
		customer[i].CannotLoginUntilDateUtc = nil

		billingAddressID, err := primitive.ObjectIDFromHex(adress[i].ID.Hex())
		if err != nil {
			result.Status = false
			result.Details = err.Error()
			return result
		}
		customer[i].BillingAddressID = &billingAddressID

		shippingAddressID, err := primitive.ObjectIDFromHex(adress[i].ID.Hex())
		if err != nil {
			result.Status = false
			result.Details = err.Error()
			return result
		}
		customer[0].ShippingAddressID = &shippingAddressID

		err = lc.InstallUsecase.UpdateCustomer(c, customer[i])
		if err != nil {
			result.Status = false
			result.Details = err.Error()
			return result
		}

		result, cusp := service.InstallCustomerPassword(customer[i].ID, model.AdminPassword)
		if !result.Status {
			return result
		}

		if len(cusp) == 0 || cusp == nil {
			result.Status = false
			result.Details = "No CustomerPassword to install"
			return result
		}

		err = lc.InstallUsecase.InstallCustomerPassword(c, cusp)
		if err != nil {
			result.Status = false
			result.Details = err.Error()
			return result
		}

		result, adressmap := service.InstallCustomerAddressMapping(customer[i].ID, adress[i].ID)
		if !result.Status {
			return result
		}

		err = lc.InstallUsecase.InstallCustomerAddressMapping(c, adressmap)
		if err != nil {
			result.Status = false
			result.Details = err.Error()
			return result
		}
	}

	return result
}

func CustomerData(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, customer := service.InstallCustomer(sample)
	if !result.Status {
		return result
	}

	if len(customer) == 0 || customer == nil {
		result.Status = false
		result.Details = "No customer to install"
		return result
	}

	err := lc.InstallUsecase.InstallCustomer(c, customer)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, custroles := service.InstallCustomerCustomerRoleMappings(sample)
	if !result.Status {
		return result
	}

	err = lc.InstallUsecase.InstallCustomerCustomerRoleMapping(c, custroles)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	///Adresses
	result, adress := service.InstallAddress(sample)
	if !result.Status {
		return result
	}

	if len(adress) == 0 || adress == nil {
		result.Status = false
		result.Details = "No Adresses to install"
		return result
	}

	err = lc.InstallUsecase.InstallAddress(c, adress)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	for i := 0; i < len(customer); i++ {

		customer[i].CreatedOnUtc = time.Now()
		customer[i].LastActivityDateUtc = time.Now()
		customer[i].LastIpAddress = c.ClientIP()
		customer[i].LastLoginDateUtc = nil
		customer[i].CannotLoginUntilDateUtc = nil

		billingAddressID, err := primitive.ObjectIDFromHex(adress[i].ID.Hex())
		if err != nil {
			result.Status = false
			result.Details = err.Error()
			return result
		}
		customer[i].BillingAddressID = &billingAddressID

		shippingAddressID, err := primitive.ObjectIDFromHex(adress[i].ID.Hex())
		if err != nil {
			result.Status = false
			result.Details = err.Error()
			return result
		}
		customer[0].ShippingAddressID = &shippingAddressID

		err = lc.InstallUsecase.UpdateCustomer(c, customer[i])
		if err != nil {
			result.Status = false
			result.Details = err.Error()
			return result
		}

		result, cusp := service.InstallCustomerPassword(customer[i].ID, customer[i].Email)
		if !result.Status {
			return result
		}

		if len(cusp) == 0 || cusp == nil {
			result.Status = false
			result.Details = "No CustomerPassword to install"
			return result
		}

		err = lc.InstallUsecase.InstallCustomerPassword(c, cusp)
		if err != nil {
			result.Status = false
			result.Details = err.Error()
			return result
		}

		result, adressmap := service.InstallCustomerAddressMapping(customer[i].ID, adress[i].ID)
		if !result.Status {
			return result
		}

		err = lc.InstallUsecase.InstallCustomerAddressMapping(c, adressmap)
		if err != nil {
			result.Status = false
			result.Details = err.Error()
			return result
		}
	}

	return result
}

func (lc *InstallController) InstallCheckoutAttribute(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := CheckoutAttribute(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CheckoutAttribute(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, items := service.InstallCheckoutAttribute(sample)

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No CheckoutAttribute to install"
		return result
	}

	err := lc.InstallUsecase.InstallCheckoutAttribute(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, values := service.InstallCheckoutAttributeValue(sample)
	if !result.Status {
		return result
	}

	err = lc.InstallUsecase.InstallCheckoutAttributeValue(c, values)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallSpecificationAttribute(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := SpecificationAttribute(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func SpecificationAttribute(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, groups := service.InstallSpecificationAttributeGroup(sample)
	if !result.Status {
		return result
	}

	err := lc.InstallUsecase.InstallSpecificationAttributeGroup(c, groups)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, items := service.InstallSpecificationAttribute(sample)
	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No SpecificationAttribute to install"
		return result
	}

	err = lc.InstallUsecase.InstallSpecificationAttribute(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, options := service.InstallSpecificationAttributeOption(sample)
	if !result.Status {
		return result
	}

	err = lc.InstallUsecase.InstallSpecificationAttributeOption(c, options)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallProductAttribute(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := ProductAttribute(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func ProductAttribute(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, items := service.InstallProductAttribute(sample)
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No ProductAttribute to install"
		return result
	}

	err := lc.InstallUsecase.InstallProductAttribute(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallCategory(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := Category(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Category(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, items := service.InstallCategory(sample)
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No Category to install"
		return result
	}

	err := lc.InstallUsecase.InstallCategory(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallPicture(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := Picture(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Picture(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, items := service.InstallPicture(sample)
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No Picture to install"
		return result
	}

	err := lc.InstallUsecase.InstallPicture(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallManufacturer(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := Manufacturer(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Manufacturer(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, items := service.InstallManufacturer(sample)
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No Manufacturer to install"
		return result
	}

	err := lc.InstallUsecase.InstallManufacturer(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallWarehouse(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := Warehouse(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Warehouse(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, items := service.InstallWarehouse(sample)
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No Warehouse to install"
		return result
	}

	err := lc.InstallUsecase.InstallWarehouse(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallVendor(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := Vendor(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Vendor(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, items := service.InstallVendor(sample)
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No Warehouse to install"
		return result
	}

	err := lc.InstallUsecase.InstallVendor(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallAffiliate(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := Affiliate(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Affiliate(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, items := service.InstallAffiliate(sample)
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No Warehouse to install"
		return result
	}

	err := lc.InstallUsecase.InstallAffiliate(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallForum(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := Forum(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Forum(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, groups := service.InstallForumGroup(sample)
	if !result.Status {
		return result
	}

	err := lc.InstallUsecase.InstallForumGroup(c, groups)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, items := service.InstallForum(sample)
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No Forum to install"
		return result
	}

	err = lc.InstallUsecase.InstallForum(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallDiscount(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := Discount(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Discount(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, items := service.InstallDiscount(sample)
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No Discount to install"
		return result
	}
	err := lc.InstallUsecase.InstallDiscount(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallBlogPost(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := BlogPost(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func BlogPost(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, items := service.InstallBlogPost(sample)
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No BlogPost to install"
		return result
	}

	err := lc.InstallUsecase.InstallBlogPost(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, comments := service.InstallBlogComment(sample)
	if !result.Status {
		return result
	}

	if len(comments) == 0 || comments == nil {
		result.Status = false
		result.Details = "No Blog commentsto install"
		return result
	}

	err = lc.InstallUsecase.InstallBlogComment(c, comments)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallPoll(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := Poll(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Poll(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, items := service.InstallPoll(sample)
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No Poll commentsto install"
		return result
	}

	err := lc.InstallUsecase.InstallPoll(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, comments := service.InstallPollAnswer(sample)
	if !result.Status {
		return result
	}

	if len(comments) == 0 || comments == nil {
		result.Status = false
		result.Details = "No PollAnswer commentsto install"
		return result
	}

	err = lc.InstallUsecase.InstallPollAnswer(c, comments)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallNewsItem(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := NewsItem(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func NewsItem(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, items := service.InstallNewsItem(sample)
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No NewsItem commentsto install"
		return result
	}

	err := lc.InstallUsecase.InstallNewsItem(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, comments := service.InstallNewsComment(sample)
	if !result.Status {
		return result
	}

	if len(comments) == 0 || comments == nil {
		result.Status = false
		result.Details = "No News Comments commentsto install"
		return result
	}

	err = lc.InstallUsecase.InstallNewsComment(c, comments)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallSearchTerm(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := SearchTerm(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func SearchTerm(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, items := service.InstallSearchTerm(sample)
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No SearchTerm commentsto install"
		return result
	}

	err := lc.InstallUsecase.InstallSearchTerm(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallProduct(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := Product(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Product(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, products := service.InstallProduct(sample)
	if !result.Status {
		return result
	}

	if len(products) == 0 || products == nil {
		result.Status = false
		result.Details = "No Products commentsto install"
		return result
	}

	for i := range products {

		timein := time.Now()
		timeout := time.Now().Add(time.Hour * 24 * 30)

		products[i].CreatedOnUtc = timein
		products[i].UpdatedOnUtc = timein

		if products[i].MarkAsNew {
			products[i].MarkAsNewStartDateTimeUtc = &timein
			products[i].MarkAsNewEndDateTimeUtc = &timeout
		}
	}

	err := lc.InstallUsecase.InstallProduct(c, products)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, caetegories := service.InstallProductCategory(sample)
	if !result.Status {
		return result
	}

	if len(caetegories) == 0 || caetegories == nil {
		result.Status = false
		result.Details = "No caetegories to install"
		return result
	}

	err = lc.InstallUsecase.InstallProductCategory(c, caetegories)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, pictures := service.InstallProductPicture(sample)
	if !result.Status {
		return result
	}

	if len(pictures) == 0 || pictures == nil {
		result.Status = false
		result.Details = "No pictures to install"
		return result
	}

	err = lc.InstallUsecase.InstallProductPicture(c, pictures)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, picturesatr := service.InstallProductAttributeValuePicture(sample)
	if !result.Status {
		return result
	}

	if len(picturesatr) == 0 || picturesatr == nil {
		result.Status = false
		result.Details = "No ProductAttributeValuePicture to install"
		return result
	}

	err = lc.InstallUsecase.InstallProductAttributeValuePicture(c, picturesatr)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, tags := service.InstallProductTag(sample)
	if !result.Status {
		return result
	}

	if len(tags) == 0 || tags == nil {
		result.Status = false
		result.Details = "No tags to install"
		return result
	}

	err = lc.InstallUsecase.InstallProductTag(c, tags)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, tagsmap := service.InstallProductProductTagMapping(sample)
	if !result.Status {
		return result
	}

	if len(tagsmap) == 0 || tagsmap == nil {
		result.Status = false
		result.Details = "No ProductTagMapping( to install"
		return result
	}

	err = lc.InstallUsecase.InstallProductProductTagMapping(c, tagsmap)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, atrmap := service.InstallProductAttributeMapping(sample)
	if !result.Status {
		return result
	}

	if len(atrmap) == 0 || atrmap == nil {
		result.Status = false
		result.Details = "No ProductAttributeMapping to install"
		return result
	}

	err = lc.InstallUsecase.InstallProductAttributeMapping(c, atrmap)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, atrvalue := service.InstallProductAttributeValue(sample)
	if !result.Status {
		return result
	}

	if len(atrvalue) == 0 || atrvalue == nil {
		result.Status = false
		result.Details = "No ProductAttributeValue to install"
		return result
	}

	err = lc.InstallUsecase.InstallProductAttributeValue(c, atrvalue)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}
	result, atrsp := service.InstallProductSpecificationAttribute(sample)
	if !result.Status {
		return result
	}

	if len(atrsp) == 0 || atrsp == nil {
		result.Status = false
		result.Details = "No ProductSpecificationAttribute to install"
		return result
	}

	err = lc.InstallUsecase.InstallProductSpecificationAttribute(c, atrsp)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, manufa := service.InstallProductManufacturer(sample)
	if !result.Status {
		return result
	}

	if len(manufa) == 0 || manufa == nil {
		result.Status = false
		result.Details = "No ProductManufacture to install"
		return result
	}

	err = lc.InstallUsecase.InstallProductManufacturer(c, manufa)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, tierprice := service.InstallTierPrice(sample)
	if !result.Status {
		return result
	}

	if len(tierprice) == 0 || tierprice == nil {
		result.Status = false
		result.Details = "No TierPrice to install"
		return result
	}

	err = lc.InstallUsecase.InstallTierPrice(c, tierprice)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, related := service.InstallRelatedProduct(sample)
	if !result.Status {
		return result
	}

	if len(related) == 0 || related == nil {
		result.Status = false
		result.Details = "No RelatedProduct to install"
		return result
	}

	err = lc.InstallUsecase.InstallRelatedProduct(c, related)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, reviews := service.InstallProductReview(sample, products)
	if !result.Status {
		return result
	}

	if len(reviews) == 0 || reviews == nil {
		result.Status = false
		result.Details = "No ProductReview to install"
		return result
	}

	err = lc.InstallUsecase.InstallProductReview(c, reviews)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, stock := service.InstallStockQuantityChange(sample, products)
	if !result.Status {
		return result
	}

	if len(stock) == 0 || stock == nil {
		result.Status = false
		result.Details = "No StockQuantityChange to install"
		return result
	}

	err = lc.InstallUsecase.InstallStockQuantityChange(c, stock)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result.Status = true
	result.Details = "Products with details data installed successfully"

	return result
}

func (lc *InstallController) InstallDownload(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := Download(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Download(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, items := service.InstallDownload(sample)
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No Download to install"
		return result
	}

	err := lc.InstallUsecase.InstallDownload(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallGdprConsent(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := GdprConsent(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func GdprConsent(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, items := service.InstallGdprConsent(false)
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No GdprConsent to install"
		return result
	}

	err := lc.InstallUsecase.InstallGdprConsent(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	return result
}

func (lc *InstallController) InstallOrder(c *gin.Context) {
	isSample := false
	sample := c.Query("sample")
	if sample == "true" {
		isSample = true
	}
	result := Order(c, lc, isSample)
	if !result.Status {
		c.JSON(http.StatusFailedDependency, common.ErrorResponse{Message: result.Details})
		return
	}
	c.JSON(http.StatusOK, result)
}

func Order(c *gin.Context, lc *InstallController, sample bool) response.Install {

	result, order := service.InstallOrder(sample)
	if !result.Status {
		return result
	}

	if len(order) == 0 || order == nil {
		result.Status = false
		result.Details = "No Order to install"
		return result
	}

	err := lc.InstallUsecase.InstallOrder(c, order)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, items := service.InstallOrderItem(sample, order)
	if !result.Status {
		return result
	}

	if len(items) == 0 || items == nil {
		result.Status = false
		result.Details = "No OrderItem to install"
		return result
	}

	err = lc.InstallUsecase.InstallOrderItem(c, items)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, ordern := service.InstallOrderNote(sample)
	if !result.Status {
		return result
	}

	if len(ordern) == 0 || ordern == nil {
		result.Status = false
		result.Details = "No OrderNote to install"
		return result
	}

	err = lc.InstallUsecase.InstallOrderNote(c, ordern)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, shipp := service.InstallShipment(sample)
	if !result.Status {
		return result
	}

	if len(shipp) == 0 || shipp == nil {
		result.Status = false
		result.Details = "No Shipment to install"
		return result
	}

	err = lc.InstallUsecase.InstallShipment(c, shipp)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result, shippit := service.InstallShipmentItem(sample)
	if !result.Status {
		return result
	}

	if len(shippit) == 0 || shippit == nil {
		result.Status = false
		result.Details = "No ShipmentItem to install"
		return result
	}

	err = lc.InstallUsecase.InstallShipmentItem(c, shippit)
	if err != nil {
		result.Status = false
		result.Details = err.Error()
		return result
	}

	result.Status = true
	result.Details = "Orders with items and shippings data installed successfully"
	return result
}
