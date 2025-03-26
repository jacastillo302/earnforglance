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

func ReturnRequestActionRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewReturnRequestActionRepository(db, domain.CollectionReturnRequestAction)
	lc := &controller.ReturnRequestActionController{
		ReturnRequestActionUsecase: usecase.NewReturnRequestActionUsecase(ur, timeout),
		Env:                        env,
	}
	itemGroup := group.Group("/api/v1/orders")
	itemGroup.GET("/return_request_actions", lc.Fetch)
	itemGroup.GET("/return_request_action", lc.FetchByID)
	itemGroup.POST("/return_request_action", lc.Create)
	itemGroup.POST("/return_request_actions", lc.CreateMany)
	itemGroup.PUT("/return_request_action", lc.Update)
	itemGroup.DELETE("/return_request_action", lc.Delete)
}
