package route

import (
	"time"

	controller "earnforglance/server/api/controller/public"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/public"

	repository "earnforglance/server/repository/public"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/public"

	"github.com/gin-gonic/gin"
)

func TopicRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewTopicRepository(db, domain.CollectionUser)
	tp := &controller.TopicController{
		TopicUsecase: usecase.NewtopicUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/topics", tp.GetTopics)
	group.POST("/topic_secret", tp.GetTopicSecret)
}
