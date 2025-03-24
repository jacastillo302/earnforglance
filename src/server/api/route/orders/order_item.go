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

func OrderItemRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewOrderItemRepository(db, domain.CollectionOrderItem)
	lc := &controller.OrderItemController{
		OrderItemUsecase: usecase.NewOrderItemUsecase(ur, timeout),
		Env:              env,
	}

	group.GET("/orderitems", lc.Fetch)
	group.GET("/orderitem", lc.FetchByID)
	group.POST("/orderitem", lc.Create)
	group.POST("/orderitems", lc.CreateMany)
	group.PUT("/orderitem", lc.Update)
	group.DELETE("/orderitem", lc.Delete)
}
