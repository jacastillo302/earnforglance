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

	group.GET("/measure_settings", lc.Fetch)
	group.GET("/measure_setting", lc.FetchByID)
	group.POST("/measure_setting", lc.Create)
	group.POST("/measure_settings", lc.CreateMany)
	group.PUT("/measure_setting", lc.Update)
	group.DELETE("/measure_setting", lc.Delete)
}
