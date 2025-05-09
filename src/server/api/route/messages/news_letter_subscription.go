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

func NewsLetterSubscriptionRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewNewsLetterSubscriptionRepository(db, domain.CollectionNewsLetterSubscription)
	lc := &controller.NewsLetterSubscriptionController{
		NewsLetterSubscriptionUsecase: usecase.NewNewsLetterSubscriptionUsecase(ur, timeout),
		Env:                           env,
	}
	itemGroup := group.Group("/api/v1/messages")
	itemGroup.GET("/news_letter_subscriptions", lc.Fetch)
	itemGroup.GET("/news_letter_subscription", lc.FetchByID)
	itemGroup.POST("/news_letter_subscription", lc.Create)
	itemGroup.POST("/news_letter_subscriptions", lc.CreateMany)
	itemGroup.PUT("/news_letter_subscription", lc.Update)
	itemGroup.DELETE("/news_letter_subscription", lc.Delete)
}
