package route

import (
	"time"

	controller "earnforglance/server/api/controller/orders"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/orders"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/orders"
	usecase "earnforglance/server/usecase/orders"

	"github.com/gin-gonic/gin"
)

func RecurringPaymentRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewRecurringPaymentRepository(db, domain.CollectionRecurringPayment)
	lc := &controller.RecurringPaymentController{
		RecurringPaymentUsecase: usecase.NewRecurringPaymentUsecase(ur, timeout),
		Env:                     env,
	}

	group.GET("/recurring_payments", lc.Fetch)
	group.GET("/recurring_payment", lc.FetchByID)
	group.POST("/recurring_payment", lc.Create)
	group.PUT("/recurring_payment", lc.Update)
	group.DELETE("/recurring_payment", lc.Delete)
}
