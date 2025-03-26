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
	itemGroup := group.Group("/api/v1/orders")
	itemGroup.GET("/recurring_payments", lc.Fetch)
	itemGroup.GET("/recurring_payment", lc.FetchByID)
	itemGroup.POST("/recurring_payment", lc.Create)
	itemGroup.POST("/recurring_payments", lc.CreateMany)
	itemGroup.PUT("/recurring_payment", lc.Update)
	itemGroup.DELETE("/recurring_payment", lc.Delete)
}
