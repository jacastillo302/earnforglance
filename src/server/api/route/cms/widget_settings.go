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

	group.GET("/widget_settings", lc.Fetch)
	group.GET("/widget_setting", lc.FetchByID)
	group.POST("/widget_setting", lc.Create)
	group.POST("/widget_settings", lc.CreateMany)
	group.PUT("/widget_setting", lc.Update)
	group.DELETE("widget_setting", lc.Delete)
}
