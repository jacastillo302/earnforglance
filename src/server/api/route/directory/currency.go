package route

import (
	"time"

	controller "earnforglance/server/api/controller/directory"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/directory"

	repository "earnforglance/server/repository/directory"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/directory"

	"github.com/gin-gonic/gin"
)

func CurrencyRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCurrencyRepository(db, domain.CollectionCurrency)
	lc := &controller.CurrencyController{
		CurrencyUsecase: usecase.NewCurrencyUsecase(ur, timeout),
		Env:             env,
	}

	itemGroup := group.Group("/api/v1/directory")
	itemGroup.GET("/currencies", lc.Fetch)
	itemGroup.GET("/currency", lc.FetchByID)
	itemGroup.POST("/currency", lc.Create)
	itemGroup.POST("/currencies", lc.CreateMany)
	itemGroup.PUT("/currency", lc.Update)
	itemGroup.DELETE("/currency", lc.Delete)
}
