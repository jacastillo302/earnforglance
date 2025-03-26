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

func ReturnRequestRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewReturnRequestRepository(db, domain.CollectionReturnRequest)
	lc := &controller.ReturnRequestController{
		ReturnRequestUsecase: usecase.NewReturnRequestUsecase(ur, timeout),
		Env:                  env,
	}
	itemGroup := group.Group("/api/v1/orders")
	itemGroup.GET("/return_requests", lc.Fetch)
	itemGroup.GET("/return_request", lc.FetchByID)
	itemGroup.POST("/return_request", lc.Create)
	itemGroup.POST("/return_requests", lc.CreateMany)
	itemGroup.PUT("/return_request", lc.Update)
	itemGroup.DELETE("/return_request", lc.Delete)
}
