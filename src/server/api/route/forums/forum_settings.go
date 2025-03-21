package route

import (
	"time"

	controller "earnforglance/server/api/controller/forums"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/forums"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/forums"
	usecase "earnforglance/server/usecase/forums"

	"github.com/gin-gonic/gin"
)

func ForumSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewForumSettingsRepository(db, domain.CollectionForumSettings)
	lc := &controller.ForumSettingsController{
		ForumSettingsUsecase: usecase.NewForumSettingsUsecase(ur, timeout),
		Env:                  env,
	}

	group.GET("/forum_settingss", lc.Fetch)
	group.GET("/forum_settings", lc.FetchByID)
	group.POST("/forum_settings", lc.Create)
	group.PUT("/forum_settings", lc.Update)
	group.DELETE("forum_settings", lc.Delete)
}
