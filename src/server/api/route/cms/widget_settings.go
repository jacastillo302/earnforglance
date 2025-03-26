package route

import (
	"time"

	controller "earnforglance/server/api/controller/cms"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/cms"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/cms"
	usecase "earnforglance/server/usecase/cms"

	"github.com/gin-gonic/gin"
)

func WidgetSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewWidgetSettingsRepository(db, domain.CollectionWidgetSettings)
	lc := &controller.WidgetSettingsController{
		WidgetSettingsUsecase: usecase.NewWidgetSettingsUsecase(ur, timeout),
		Env:                   env,
	}

	itemGroup := group.Group("/api/v1/cms")
	itemGroup.GET("/widget_settings", lc.Fetch)
	itemGroup.GET("/widget_setting", lc.FetchByID)
	itemGroup.POST("/widget_setting", lc.Create)
	itemGroup.POST("/widget_settings", lc.CreateMany)
	itemGroup.PUT("/widget_setting", lc.Update)
	itemGroup.DELETE("/widget_setting", lc.Delete)
}
