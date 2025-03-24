package route

import (
	"time"

	controller "earnforglance/server/api/controller/messages"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/messages"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/messages"
	usecase "earnforglance/server/usecase/messages"

	"github.com/gin-gonic/gin"
)

func NewsLetterSubscriptionRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewNewsLetterSubscriptionRepository(db, domain.CollectionNewsLetterSubscription)
	lc := &controller.NewsLetterSubscriptionController{
		NewsLetterSubscriptionUsecase: usecase.NewNewsLetterSubscriptionUsecase(ur, timeout),
		Env:                           env,
	}

	group.GET("/news_letter_subscriptions", lc.Fetch)
	group.GET("/news_letter_subscription", lc.FetchByID)
	group.POST("/news_letter_subscription", lc.Create)
	group.POST("/news_letter_subscriptions", lc.CreateMany)
	group.PUT("/news_letter_subscription", lc.Update)
	group.DELETE("/news_letter_subscription", lc.Delete)
}
