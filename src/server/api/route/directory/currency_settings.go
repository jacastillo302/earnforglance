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

func CurrencySettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCurrencySettingsRepository(db, domain.CollectionCurrencySettings)
	lc := &controller.CurrencySettingsController{
		CurrencySettingsUsecase: usecase.NewCurrencySettingsUsecase(ur, timeout),
		Env:                     env,
	}

	group.GET("/currency_settings", lc.Fetch)
	group.GET("/currency_setting", lc.FetchByID)
	group.POST("/currency_setting", lc.Create)
	group.PUT("/currency_setting", lc.Update)
	group.DELETE("/currency_setting", lc.Delete)
}
