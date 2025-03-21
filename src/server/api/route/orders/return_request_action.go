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

func ReturnRequestActionRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewReturnRequestActionRepository(db, domain.CollectionReturnRequestAction)
	lc := &controller.ReturnRequestActionController{
		ReturnRequestActionUsecase: usecase.NewReturnRequestActionUsecase(ur, timeout),
		Env:                        env,
	}

	group.GET("/return_request_actions", lc.Fetch)
	group.GET("/return_request_action", lc.FetchByID)
	group.POST("/return_request_action", lc.Create)
	group.PUT("/return_request_action", lc.Update)
	group.DELETE("/return_request_action", lc.Delete)
}
