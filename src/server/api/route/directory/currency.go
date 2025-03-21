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

func CurrencyRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCurrencyRepository(db, domain.CollectionCurrency)
	lc := &controller.CurrencyController{
		CurrencyUsecase: usecase.NewCurrencyUsecase(ur, timeout),
		Env:             env,
	}

	group.GET("/currencies", lc.Fetch)
	group.GET("/currency", lc.FetchByID)
	group.POST("/currency", lc.Create)
	group.PUT("/currency", lc.Update)
	group.DELETE("/currency", lc.Delete)
}
