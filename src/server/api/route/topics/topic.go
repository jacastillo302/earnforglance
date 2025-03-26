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

func TopicRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewTopicRepository(db, domain.CollectionTopic)
	lc := &controller.TopicController{
		TopicUsecase: usecase.NewTopicUsecase(ur, timeout),
		Env:          env,
	}
	itemGroup := group.Group("/api/v1/topics")
	itemGroup.GET("/topics", lc.Fetch)
	itemGroup.GET("/topic", lc.FetchByID)
	itemGroup.POST("/topic", lc.Create)
	itemGroup.POST("/topics", lc.CreateMany)
	itemGroup.PUT("/topic", lc.Update)
	itemGroup.DELETE("/topic", lc.Delete)
}
