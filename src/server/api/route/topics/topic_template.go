package route

import (
	"time"

	controller "earnforglance/server/api/controller/topics"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/topics"

	repository "earnforglance/server/repository/topics"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/topics"

	"github.com/gin-gonic/gin"
)

func TopicTemplateRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewTopicTemplateRepository(db, domain.CollectionTopicTemplate)
	lc := &controller.TopicTemplateController{
		TopicTemplateUsecase: usecase.NewTopicTemplateUsecase(ur, timeout),
		Env:                  env,
	}
	itemGroup := group.Group("/api/v1/topics")
	itemGroup.GET("/topic_templates", lc.Fetch)
	itemGroup.GET("/topic_template", lc.FetchByID)
	itemGroup.POST("/topic_template", lc.Create)
	itemGroup.POST("/topic_templates", lc.CreateMany)
	itemGroup.PUT("/topic_template", lc.Update)
	itemGroup.DELETE("/topic_template", lc.Delete)
}
