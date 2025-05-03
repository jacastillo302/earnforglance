package route

import (
	"context"
	controller "earnforglance/server/api/controller/public"
	"earnforglance/server/bootstrap"
	customer "earnforglance/server/domain/customers"
	news "earnforglance/server/domain/news"
	domain "earnforglance/server/domain/public"
	"time"

	repository "earnforglance/server/repository/public"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/public"

	"github.com/gin-gonic/gin"
)

func CustomerRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	nr := repository.NewNewsLetterRepository(db, news.CollectionNewsItem)
	ur := repository.NewCustomerRepository(db, domain.CollectionUser, nr)
	cc := &controller.CustomerController{
		CustomerUsecase: usecase.NewCustomerUsecase(ur, nr, timeout),
		Env:             env,
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	slugs, _ := cc.CustomerUsecase.GetSlugs(ctx, customer.CollectionCustomer)
	for _, slug := range slugs {
		group.POST(slug, cc.SignIn)
	}
}
