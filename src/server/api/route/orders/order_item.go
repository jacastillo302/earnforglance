package route

import (
	"time"

	controller "earnforglance/server/api/controller/orders"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/orders"

	repository "earnforglance/server/repository/orders"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/orders"

	"github.com/gin-gonic/gin"
)

func OrderItemRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewOrderItemRepository(db, domain.CollectionOrderItem)
	lc := &controller.OrderItemController{
		OrderItemUsecase: usecase.NewOrderItemUsecase(ur, timeout),
		Env:              env,
	}
	itemGroup := group.Group("/api/v1/orders")
	itemGroup.GET("/order_items", lc.Fetch)
	itemGroup.GET("/order_item", lc.FetchByID)
	itemGroup.POST("/order_item", lc.Create)
	itemGroup.POST("/order_items", lc.CreateMany)
	itemGroup.PUT("/order_item", lc.Update)
	itemGroup.DELETE("/order_item", lc.Delete)
}
