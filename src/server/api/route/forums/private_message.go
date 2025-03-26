package route

import (
	"time"

	controller "earnforglance/server/api/controller/forums"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/forums"

	repository "earnforglance/server/repository/forums"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/forums"

	"github.com/gin-gonic/gin"
)

func PrivateMessageRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPrivateMessageRepository(db, domain.CollectionPrivateMessage)
	lc := &controller.PrivateMessageController{
		PrivateMessageUsecase: usecase.NewPrivateMessageUsecase(ur, timeout),
		Env:                   env,
	}
	itemGroup := group.Group("/api/v1/forums")
	itemGroup.GET("/private_messages", lc.Fetch)
	itemGroup.GET("/private_message", lc.FetchByID)
	itemGroup.POST("/private_message", lc.Create)
	itemGroup.POST("/private_messages", lc.CreateMany)
	itemGroup.PUT("/private_message", lc.Update)
	itemGroup.DELETE("/private_message", lc.Delete)
}
