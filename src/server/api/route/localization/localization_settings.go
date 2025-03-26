package route

import (
	"time"

	controller "earnforglance/server/api/controller/localization"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/localization"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/localization"
	usecase "earnforglance/server/usecase/localization"

	"github.com/gin-gonic/gin"
)

func LocalizationSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewLocalizationSettingsRepository(db, domain.CollectionLocalizationSettings)
	lc := &controller.LocalizationSettingsController{
		LocalizationSettingsUsecase: usecase.NewLocalizationSettingsUsecase(ur, timeout),
		Env:                         env,
	}
	itemGroup := group.Group("/api/v1/localization")
	itemGroup.GET("/localization_settings", lc.Fetch)
	itemGroup.GET("/localization_setting", lc.FetchByID)
	itemGroup.POST("/localization_setting", lc.Create)
	itemGroup.POST("/localization_settings", lc.CreateMany)
	itemGroup.PUT("/localization_setting", lc.Update)
	itemGroup.DELETE("/localization_setting", lc.Delete)
}
