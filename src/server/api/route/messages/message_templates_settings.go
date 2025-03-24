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

func MessageTemplatesSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewMessageTemplatesSettingsRepository(db, domain.CollectionMessageTemplatesSettings)
	lc := &controller.MessageTemplatesSettingsController{
		MessageTemplatesSettingsUsecase: usecase.NewMessageTemplatesSettingsUsecase(ur, timeout),
		Env:                             env,
	}

	group.GET("/message_templates_settings", lc.Fetch)
	group.GET("/message_templates_setting", lc.FetchByID)
	group.POST("/message_templates_setting", lc.Create)
	group.POST("/message_templates_settings", lc.CreateMany)
	group.PUT("/message_templates_setting", lc.Update)
	group.DELETE("message_templates_setting", lc.Delete)
}
