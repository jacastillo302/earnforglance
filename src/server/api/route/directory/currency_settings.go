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

	itemGroup := group.Group("/api/v1/directory")
	itemGroup.GET("/currency_settings", lc.Fetch)
	itemGroup.GET("/currency_setting", lc.FetchByID)
	itemGroup.POST("/currency_setting", lc.Create)
	itemGroup.POST("/currency_settings", lc.CreateMany)
	itemGroup.PUT("/currency_setting", lc.Update)
	itemGroup.DELETE("/currency_setting", lc.Delete)
}
