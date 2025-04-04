package route

import (
	"time"

	controller "earnforglance/server/api/controller/install"
	"earnforglance/server/bootstrap"

	repository "earnforglance/server/repository/install"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/install"

	"github.com/gin-gonic/gin"
)

func InstallRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewInstallRepository(db)
	lc := &controller.InstallController{
		InstallUsecase: usecase.NewInstallUsecase(ur, timeout),
		Env:            env,
	}

	group.GET("/ping_database", lc.PingDatabase)
	group.GET("/permission_record", lc.InstallPermissionRecord)
	group.GET("/currency", lc.InstallCurrencies)
	group.GET("/measure_dimension", lc.InstallMeasureDimension)
	group.GET("/measure_weight", lc.InstallMeasureWeight)
	group.GET("/tax_category", lc.InstallTaxCategories)
	group.GET("/language", lc.InstallLanguages)
	group.GET("/store", lc.InstallStores)
	group.GET("/setting", lc.InstallSettings)
	group.GET("/country", lc.InstallCountries)
	group.GET("/shipping_method", lc.InstallShippingMethod)
	group.GET("/delivery_date", lc.InstallDeliveryDate)
	group.GET("/product_availability_range", lc.InstallProductAvailabilityRange)
	group.GET("/email_account", lc.InstallEmailAccount)
	group.GET("/message_template", lc.InstallMessageTemplate)
	group.GET("/topic_template", lc.InstallTopicTemplate)
	group.GET("/customer", lc.InstallCustomerRole)
	group.GET("/activity_log_type", lc.InstallActivityLogType)
	group.GET("/product_template", lc.InstallProductTemplate)
	group.GET("/category_template", lc.InstallCategoryTemplate)
	group.GET("/manufacturer_template", lc.InstallManufacturerTemplate)
	group.GET("/schedule_task", lc.InstallScheduleTask)
	group.GET("/return_request_reason", lc.InstallReturnRequestReason)
	group.GET("/return_request_action", lc.InstallReturnRequestAction)

}
