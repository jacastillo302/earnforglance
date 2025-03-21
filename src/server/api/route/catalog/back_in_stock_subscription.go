package route

import (
	"time"

	controller "earnforglance/server/api/controller/catalog"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/catalog"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/catalog"
	usecase "earnforglance/server/usecase/catalog"

	"github.com/gin-gonic/gin"
)

func BackInStockSubscriptionRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewBackInStockSubscriptionRepository(db, domain.CollectionBackInStockSubscription)
	lc := &controller.BackInStockSubscriptionController{
		BackInStockSubscriptionUsecase: usecase.NewBackInStockSubscriptionUsecase(ur, timeout),
		Env:                            env,
	}

	group.GET("/back_in_stock_subscriptions", lc.Fetch)
	group.GET("/back_in_stock_subscription", lc.FetchByID)
	group.POST("/back_in_stock_subscription", lc.Create)
	group.PUT("/back_in_stock_subscription", lc.Update)
	group.DELETE("/back_in_stock_subscription", lc.Delete)
}
