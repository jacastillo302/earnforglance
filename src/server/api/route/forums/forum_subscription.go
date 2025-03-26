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

func ForumSubscriptionRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewForumSubscriptionRepository(db, domain.CollectionForumSubscription)
	lc := &controller.ForumSubscriptionController{
		ForumSubscriptionUsecase: usecase.NewForumSubscriptionUsecase(ur, timeout),
		Env:                      env,
	}

	itemGroup := group.Group("/api/v1/forums")
	itemGroup.GET("/forum_subscriptions", lc.Fetch)
	itemGroup.GET("/forum_subscription", lc.FetchByID)
	itemGroup.POST("/forum_subscription", lc.Create)
	itemGroup.POST("/forum_subscriptions", lc.CreateMany)
	itemGroup.PUT("/forum_subscription", lc.Update)
	itemGroup.DELETE("/forum_subscription", lc.Delete)
}
