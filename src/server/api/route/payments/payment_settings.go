package route

import (
	"time"

	controller "earnforglance/server/api/controller/payments"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/payments"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/payments"
	usecase "earnforglance/server/usecase/payments"

	"github.com/gin-gonic/gin"
)

func PaymentSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPaymentSettingsRepository(db, domain.CollectionPaymentSettings)
	lc := &controller.PaymentSettingsController{
		PaymentSettingsUsecase: usecase.NewPaymentSettingsUsecase(ur, timeout),
		Env:                    env,
	}

	group.GET("/payment_settings", lc.Fetch)
	group.GET("/payment_setting", lc.FetchByID)
	group.POST("/payment_setting", lc.Create)
	group.PUT("/payment_setting", lc.Update)
	group.DELETE("/payment_setting", lc.Delete)
}
