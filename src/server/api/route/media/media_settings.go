package route

import (
	"time"

	controller "earnforglance/server/api/controller/media"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/media"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/media"
	usecase "earnforglance/server/usecase/media"

	"github.com/gin-gonic/gin"
)

func MediaSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewMediaSettingsRepository(db, domain.CollectionMediaSettings)
	lc := &controller.MediaSettingsController{
		MediaSettingsUsecase: usecase.NewMediaSettingsUsecase(ur, timeout),
		Env:                  env,
	}

	group.GET("/media_settings", lc.Fetch)
	group.GET("/media_setting", lc.FetchByID)
	group.POST("/media_setting", lc.Create)
	group.POST("/media_settings", lc.CreateMany)
	group.PUT("/media_setting", lc.Update)
	group.DELETE("/media_setting", lc.Delete)
}
