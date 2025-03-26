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

func OrderRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewOrderRepository(db, domain.CollectionOrder)
	lc := &controller.OrderController{
		OrderUsecase: usecase.NewOrderUsecase(ur, timeout),
		Env:          env,
	}
	itemGroup := group.Group("/api/v1/orders")
	itemGroup.GET("/orders", lc.Fetch)
	itemGroup.GET("/order", lc.FetchByID)
	itemGroup.POST("/order", lc.Create)
	itemGroup.POST("/orders", lc.CreateMany)
	itemGroup.PUT("/order", lc.Update)
	itemGroup.DELETE("/order", lc.Delete)
}
