package route

import (
	"time"

	controller "earnforglance/server/api/controller/topics"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/topics"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/topics"
	usecase "earnforglance/server/usecase/topics"

	"github.com/gin-gonic/gin"
)

func TopicTemplateRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewTopicTemplateRepository(db, domain.CollectionTopicTemplate)
	lc := &controller.TopicTemplateController{
		TopicTemplateUsecase: usecase.NewTopicTemplateUsecase(ur, timeout),
		Env:                  env,
	}

	group.GET("/topic_templates", lc.Fetch)
	group.GET("/topic_template", lc.FetchByID)
	group.POST("/topic_template", lc.Create)
	group.POST("/topic_templates", lc.CreateMany)
	group.PUT("/topic_template", lc.Update)
	group.DELETE("/topic_template", lc.Delete)
}
