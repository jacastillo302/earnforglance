package route

import (
	"time"

	controller "earnforglance/server/api/controller/catalog"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/catalog"

	repository "earnforglance/server/repository/catalog"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/catalog"

	"github.com/gin-gonic/gin"
)

func BackInStockSubscriptionRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewBackInStockSubscriptionRepository(db, domain.CollectionBackInStockSubscription)
	lc := &controller.BackInStockSubscriptionController{
		BackInStockSubscriptionUsecase: usecase.NewBackInStockSubscriptionUsecase(ur, timeout),
		Env:                            env,
	}

	itemGroup := group.Group("/api/v1/catalog")

	itemGroup.GET("/back_in_stock_subscriptions", lc.Fetch)
	itemGroup.GET("/back_in_stock_subscription", lc.FetchByID)
	itemGroup.POST("/back_in_stock_subscription", lc.Create)
	itemGroup.POST("/back_in_stock_subscriptions", lc.CreateMany)
	itemGroup.PUT("/back_in_stock_subscription", lc.Update)
	itemGroup.DELETE("/back_in_stock_subscription", lc.Delete)
}
