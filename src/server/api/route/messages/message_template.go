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

func MessageTemplateRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewMessageTemplateRepository(db, domain.CollectionMessageTemplate)
	lc := &controller.MessageTemplateController{
		MessageTemplateUsecase: usecase.NewMessageTemplateUsecase(ur, timeout),
		Env:                    env,
	}
	itemGroup := group.Group("/api/v1/messages")
	itemGroup.GET("/message_templates", lc.Fetch)
	itemGroup.GET("/message_template", lc.FetchByID)
	itemGroup.POST("/message_template", lc.Create)
	itemGroup.POST("/message_templates", lc.CreateMany)
	itemGroup.PUT("/message_template", lc.Update)
	itemGroup.DELETE("/message_template", lc.Delete)
}
