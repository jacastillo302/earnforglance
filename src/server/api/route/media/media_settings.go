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
	itemGroup := group.Group("/api/v1/media")
	itemGroup.GET("/media_settings", lc.Fetch)
	itemGroup.GET("/media_setting", lc.FetchByID)
	itemGroup.POST("/media_setting", lc.Create)
	itemGroup.POST("/media_settings", lc.CreateMany)
	itemGroup.PUT("/media_setting", lc.Update)
	itemGroup.DELETE("/media_setting", lc.Delete)
}
