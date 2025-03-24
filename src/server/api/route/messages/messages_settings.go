package route

import (
	"time"

	controller "earnforglance/server/api/controller/messages"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/messages"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/messages"
	usecase "earnforglance/server/usecase/messages"

	"github.com/gin-gonic/gin"
)

func MessagesSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewMessagesSettingsRepository(db, domain.CollectionMessagesSettings)
	lc := &controller.MessagesSettingsController{
		MessagesSettingsUsecase: usecase.NewMessagesSettingsUsecase(ur, timeout),
		Env:                     env,
	}

	group.GET("/messages_settings", lc.Fetch)
	group.GET("/messages_setting", lc.FetchByID)
	group.POST("/messages_setting", lc.Create)
	group.POST("/messages_settings", lc.CreateMany)
	group.PUT("/messages_setting", lc.Update)
	group.DELETE("/messages_setting", lc.Delete)
}
