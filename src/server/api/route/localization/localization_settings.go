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

	group.GET("/localization_settings", lc.Fetch)
	group.GET("/localization_setting", lc.FetchByID)
	group.POST("/localization_setting", lc.Create)
	group.POST("/localization_settings", lc.CreateMany)
	group.PUT("/localization_setting", lc.Update)
	group.DELETE("localization_setting", lc.Delete)
}
