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
	itemGroup.GET("/orderitems", lc.Fetch)
	itemGroup.GET("/orderitem", lc.FetchByID)
	itemGroup.POST("/orderitem", lc.Create)
	itemGroup.POST("/orderitems", lc.CreateMany)
	itemGroup.PUT("/orderitem", lc.Update)
	itemGroup.DELETE("/orderitem", lc.Delete)
}
