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

func MeasureSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewMeasureSettingsRepository(db, domain.CollectionMeasureSettings)
	lc := &controller.MeasureSettingsController{
		MeasureSettingsUsecase: usecase.NewMeasureSettingsUsecase(ur, timeout),
		Env:                    env,
	}

	group.GET("/measure_settingss", lc.Fetch)
	group.GET("/measure_settings", lc.FetchByID)
	group.POST("/measure_settings", lc.Create)
	group.PUT("/measure_settings", lc.Update)
	group.DELETE("/measure_settings", lc.Delete)
}
