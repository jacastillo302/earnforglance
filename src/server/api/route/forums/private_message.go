package route

import (
	"time"

	controller "earnforglance/server/api/controller/forums"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/forums"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/forums"
	usecase "earnforglance/server/usecase/forums"

	"github.com/gin-gonic/gin"
)

func PrivateMessageRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPrivateMessageRepository(db, domain.CollectionPrivateMessage)
	lc := &controller.PrivateMessageController{
		PrivateMessageUsecase: usecase.NewPrivateMessageUsecase(ur, timeout),
		Env:                   env,
	}

	group.GET("/private_messages", lc.Fetch)
	group.GET("/private_message", lc.FetchByID)
	group.POST("/private_message", lc.Create)
	group.POST("/private_messages", lc.CreateMany)
	group.PUT("/private_message", lc.Update)
	group.DELETE("/private_message", lc.Delete)
}
