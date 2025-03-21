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

func MessageTemplateRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewMessageTemplateRepository(db, domain.CollectionMessageTemplate)
	lc := &controller.MessageTemplateController{
		MessageTemplateUsecase: usecase.NewMessageTemplateUsecase(ur, timeout),
		Env:                    env,
	}

	group.GET("/message_templates", lc.Fetch)
	group.GET("/message_template", lc.FetchByID)
	group.POST("/message_template", lc.Create)
	group.PUT("/message_template", lc.Update)
	group.DELETE("/message_template", lc.Delete)
}
