package route

import (
	"context"
	controller "earnforglance/server/api/controller/public"
	"earnforglance/server/bootstrap"
	messages "earnforglance/server/domain/messages"
	domain "earnforglance/server/domain/public"
	"time"

	repository "earnforglance/server/repository/public"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/public"

	"github.com/gin-gonic/gin"
)

func NewsLetterRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewNewsLetterRepository(db, domain.CollectionUser)
	tp := &controller.NewsLetterController{
		NewsLetterUsecase: usecase.NewNewsLetterUsecase(ur, timeout),
		Env:               env,
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	slugs, _ := tp.NewsLetterUsecase.GetSlugs(ctx, messages.CollectionNewsLetterSubscription)

	for _, slug := range slugs {
		group.GET(slug, tp.NewsLetterSubscription)
		group.GET(slug+"_confirm", tp.NewsLetterActivation)
		group.GET(slug+"_unsuscribe", tp.NewsLetterUnSubscribe)
		group.GET(slug+"_inactivate", tp.NewsLetterInactivate)

	}

}
