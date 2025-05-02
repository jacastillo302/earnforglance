package route

import (
	"context"
	controller "earnforglance/server/api/controller/public"
	"earnforglance/server/bootstrap"
	customer "earnforglance/server/domain/customers"
	domain "earnforglance/server/domain/public"
	"time"

	repository "earnforglance/server/repository/public"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/public"

	"github.com/gin-gonic/gin"
)

func CustomerRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCustomerRepository(db, domain.CollectionUser)
	tp := &controller.CustomerController{
		CustomerUsecase: usecase.NewCustomerUsecase(ur, timeout),
		Env:             env,
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	slugs, _ := tp.CustomerUsecase.GetSlugs(ctx, customer.CollectionCustomer)
	for _, slug := range slugs {
		group.POST(slug, tp.SignIn)
	}
}
