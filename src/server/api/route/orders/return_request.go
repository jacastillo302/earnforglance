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

	group.GET("/return_requests", lc.Fetch)
	group.GET("/return_request", lc.FetchByID)
	group.POST("/return_request", lc.Create)
	group.POST("/return_requests", lc.CreateMany)
	group.PUT("/return_request", lc.Update)
	group.DELETE("/return_request", lc.Delete)
}
