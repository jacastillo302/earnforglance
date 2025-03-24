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

	group.GET("/forum_subscriptions", lc.Fetch)
	group.GET("/forum_subscription", lc.FetchByID)
	group.POST("/forum_subscription", lc.Create)
	group.POST("/forum_subscriptions", lc.CreateMany)
	group.PUT("/forum_subscription", lc.Update)
	group.DELETE("/forum_subscription", lc.Delete)
}
