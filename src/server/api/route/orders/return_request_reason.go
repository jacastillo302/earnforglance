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

func ReturnRequestReasonRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewReturnRequestReasonRepository(db, domain.CollectionReturnRequestReason)
	lc := &controller.ReturnRequestReasonController{
		ReturnRequestReasonUsecase: usecase.NewReturnRequestReasonUsecase(ur, timeout),
		Env:                        env,
	}
	itemGroup := group.Group("/api/v1/orders")
	itemGroup.GET("/return_request_reasons", lc.Fetch)
	itemGroup.GET("/return_request_reason", lc.FetchByID)
	itemGroup.POST("/return_request_reason", lc.Create)
	itemGroup.POST("/return_request_reasons", lc.CreateMany)
	itemGroup.PUT("/return_request_reason", lc.Update)
	itemGroup.DELETE("/return_request_reason", lc.Delete)
}
