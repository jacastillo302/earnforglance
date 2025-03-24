package route

import (
	"time"

	controller "earnforglance/server/api/controller/forums"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/forums"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/forums"
	usecase "earnforglance/server/usecase/forums"

	"github.com/gin-gonic/gin"
)

func ForumTopicRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewForumTopicRepository(db, domain.CollectionForumTopic)
	lc := &controller.ForumTopicController{
		ForumTopicUsecase: usecase.NewForumTopicUsecase(ur, timeout),
		Env:               env,
	}

	group.GET("/forum_topics", lc.Fetch)
	group.GET("/forum_topic", lc.FetchByID)
	group.POST("/forum_topic", lc.Create)
	group.POST("/forum_topics", lc.CreateMany)
	group.PUT("/forum_topic", lc.Update)
	group.DELETE("/forum_topic", lc.Delete)
}
