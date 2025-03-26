package route

import (
	"time"

	controller "earnforglance/server/api/controller/forums"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/forums"

	repository "earnforglance/server/repository/forums"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/forums"

	"github.com/gin-gonic/gin"
)

func ForumSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewForumSettingsRepository(db, domain.CollectionForumSettings)
	lc := &controller.ForumSettingsController{
		ForumSettingsUsecase: usecase.NewForumSettingsUsecase(ur, timeout),
		Env:                  env,
	}
	itemGroup := group.Group("/api/v1/forums")

	itemGroup.GET("/forum_settings", lc.Fetch)
	itemGroup.GET("/forum_setting", lc.FetchByID)
	itemGroup.POST("/forum_setting", lc.Create)
	itemGroup.POST("/forum_settings", lc.CreateMany)
	itemGroup.PUT("/forum_setting", lc.Update)
	itemGroup.DELETE("forum_setting", lc.Delete)
}
