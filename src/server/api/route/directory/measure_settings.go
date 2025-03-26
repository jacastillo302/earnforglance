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

func MeasureSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewMeasureSettingsRepository(db, domain.CollectionMeasureSettings)
	lc := &controller.MeasureSettingsController{
		MeasureSettingsUsecase: usecase.NewMeasureSettingsUsecase(ur, timeout),
		Env:                    env,
	}

	itemGroup := group.Group("/api/v1/directory")
	itemGroup.GET("/measure_settings", lc.Fetch)
	itemGroup.GET("/measure_setting", lc.FetchByID)
	itemGroup.POST("/measure_setting", lc.Create)
	itemGroup.POST("/measure_settings", lc.CreateMany)
	itemGroup.PUT("/measure_setting", lc.Update)
	itemGroup.DELETE("/measure_setting", lc.Delete)
}
