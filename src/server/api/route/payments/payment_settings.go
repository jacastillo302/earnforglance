package route

import (
	"time"

	controller "earnforglance/server/api/controller/payments"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/payments"

	repository "earnforglance/server/repository/payments"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/payments"

	"github.com/gin-gonic/gin"
)

func PaymentSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPaymentSettingsRepository(db, domain.CollectionPaymentSettings)
	lc := &controller.PaymentSettingsController{
		PaymentSettingsUsecase: usecase.NewPaymentSettingsUsecase(ur, timeout),
		Env:                    env,
	}
	itemGroup := group.Group("/api/v1/payments")
	itemGroup.GET("/payment_settings", lc.Fetch)
	itemGroup.GET("/payment_setting", lc.FetchByID)
	itemGroup.POST("/payment_setting", lc.Create)
	itemGroup.POST("/payment_settings", lc.CreateMany)
	itemGroup.PUT("/payment_setting", lc.Update)
	itemGroup.DELETE("/payment_setting", lc.Delete)
}
