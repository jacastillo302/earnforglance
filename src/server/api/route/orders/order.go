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

func OrderRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewOrderRepository(db, domain.CollectionOrder)
	lc := &controller.OrderController{
		OrderUsecase: usecase.NewOrderUsecase(ur, timeout),
		Env:          env,
	}

	group.GET("/orders", lc.Fetch)
	group.GET("/order", lc.FetchByID)
	group.POST("/order", lc.Create)
	group.POST("/orders", lc.CreateMany)
	group.PUT("/order", lc.Update)
	group.DELETE("/order", lc.Delete)
}
