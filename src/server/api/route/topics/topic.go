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

func TopicRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewTopicRepository(db, domain.CollectionTopic)
	lc := &controller.TopicController{
		TopicUsecase: usecase.NewTopicUsecase(ur, timeout),
		Env:          env,
	}

	group.GET("/topics", lc.Fetch)
	group.GET("/topic", lc.FetchByID)
	group.POST("/topic", lc.Create)
	group.PUT("/topic", lc.Update)
	group.DELETE("/topic", lc.Delete)
}
