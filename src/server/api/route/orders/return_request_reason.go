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

	group.GET("/return_request_reasons", lc.Fetch)
	group.GET("/return_request_reason", lc.FetchByID)
	group.POST("/return_request_reason", lc.Create)
	group.PUT("/return_request_reason", lc.Update)
	group.DELETE("/return_request_reason", lc.Delete)
}
