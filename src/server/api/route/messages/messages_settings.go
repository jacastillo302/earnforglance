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
	itemGroup := group.Group("/api/v1/messages")
	itemGroup.GET("/messages_settings", lc.Fetch)
	itemGroup.GET("/messages_setting", lc.FetchByID)
	itemGroup.POST("/messages_setting", lc.Create)
	itemGroup.POST("/messages_settings", lc.CreateMany)
	itemGroup.PUT("/messages_setting", lc.Update)
	itemGroup.DELETE("/messages_setting", lc.Delete)
}
