package route

import (
	"time"

	controller "earnforglance/server/api/controller/directory"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/directory"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/directory"
	usecase "earnforglance/server/usecase/directory"

	"github.com/gin-gonic/gin"
)

func ExchangeRateRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewExchangeRateRepository(db, domain.CollectionExchangeRate)
	lc := &controller.ExchangeRateController{
		ExchangeRateUsecase: usecase.NewExchangeRateUsecase(ur, timeout),
		Env:                 env,
	}

	itemGroup := group.Group("/api/v1/directory")
	itemGroup.GET("/exchange_rates", lc.Fetch)
	itemGroup.GET("/exchange_rate", lc.FetchByID)
	itemGroup.POST("/exchange_rate", lc.Create)
	itemGroup.POST("/exchange_rates", lc.CreateMany)
	itemGroup.PUT("/exchange_rate", lc.Update)
	itemGroup.DELETE("/exchange_rate", lc.Delete)
}
