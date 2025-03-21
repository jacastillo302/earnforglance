package route

import (
	"time"

	controller "earnforglance/server/api/controller/tax"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/tax"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/tax"
	usecase "earnforglance/server/usecase/tax"

	"github.com/gin-gonic/gin"
)

func TaxSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewTaxSettingsRepository(db, domain.CollectionTaxSettings)
	lc := &controller.TaxSettingsController{
		TaxSettingsUsecase: usecase.NewTaxSettingsUsecase(ur, timeout),
		Env:                env,
	}

	group.GET("/tax_settingss", lc.Fetch)
	group.GET("/tax_settings", lc.FetchByID)
	group.POST("/tax_settings", lc.Create)
	group.PUT("/tax_settings", lc.Update)
	group.DELETE("tax_settings", lc.Delete)
}
