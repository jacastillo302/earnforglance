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
	itemGroup := group.Group("/api/v1/tax")
	itemGroup.GET("/tax_settings", lc.Fetch)
	itemGroup.GET("/tax_setting", lc.FetchByID)
	itemGroup.POST("/tax_setting", lc.Create)
	itemGroup.POST("/tax_settings", lc.CreateMany)
	itemGroup.PUT("/tax_setting", lc.Update)
	itemGroup.DELETE("/tax_setting", lc.Delete)
}
