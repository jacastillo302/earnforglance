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

func RecurringPaymentHistoryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewRecurringPaymentHistoryRepository(db, domain.CollectionRecurringPaymentHistory)
	lc := &controller.RecurringPaymentHistoryController{
		RecurringPaymentHistoryUsecase: usecase.NewRecurringPaymentHistoryUsecase(ur, timeout),
		Env:                            env,
	}

	group.GET("/recurring_payment_histories", lc.Fetch)
	group.GET("/recurring_payment_history", lc.FetchByID)
	group.POST("/recurring_payment_history", lc.Create)
	group.POST("/recurring_payment_histories", lc.CreateMany)
	group.PUT("/recurring_payment_history", lc.Update)
	group.DELETE("/recurring_payment_history", lc.Delete)
}
