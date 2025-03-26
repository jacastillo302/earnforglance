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

	itemGroup := group.Group("/api/v1/forums")
	itemGroup.GET("/forum_topics", lc.Fetch)
	itemGroup.GET("/forum_topic", lc.FetchByID)
	itemGroup.POST("/forum_topic", lc.Create)
	itemGroup.POST("/forum_topics", lc.CreateMany)
	itemGroup.PUT("/forum_topic", lc.Update)
	itemGroup.DELETE("/forum_topic", lc.Delete)
}
