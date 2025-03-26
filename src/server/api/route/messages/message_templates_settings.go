package route

import (
	"time"

	controller "earnforglance/server/api/controller/messages"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/messages"

	repository "earnforglance/server/repository/messages"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/messages"

	"github.com/gin-gonic/gin"
)

func MessageTemplatesSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewMessageTemplatesSettingsRepository(db, domain.CollectionMessageTemplatesSettings)
	lc := &controller.MessageTemplatesSettingsController{
		MessageTemplatesSettingsUsecase: usecase.NewMessageTemplatesSettingsUsecase(ur, timeout),
		Env:                             env,
	}
	itemGroup := group.Group("/api/v1/messages")
	itemGroup.GET("/message_templates_settings", lc.Fetch)
	itemGroup.GET("/message_templates_setting", lc.FetchByID)
	itemGroup.POST("/message_templates_setting", lc.Create)
	itemGroup.POST("/message_templates_settings", lc.CreateMany)
	itemGroup.PUT("/message_templates_setting", lc.Update)
	itemGroup.DELETE("message_templates_setting", lc.Delete)
}
