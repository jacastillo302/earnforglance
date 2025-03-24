package route

import (
	"time"

	controller "earnforglance/server/api/controller/common"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/common"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/common"
	usecase "earnforglance/server/usecase/common"

	"github.com/gin-gonic/gin"
)

func DisplayDefaultMenuItemSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewDisplayDefaultMenuItemSettingsRepository(db, domain.CollectionDisplayDefaultMenuItemSettings)
	lc := &controller.DisplayDefaultMenuItemSettingsController{
		DisplayDefaultMenuItemSettingsUsecase: usecase.NewDisplayDefaultMenuItemSettingsUsecase(ur, timeout),
		Env:                                   env,
	}

	group.GET("/display_default_menu_item_settings", lc.Fetch)
	group.GET("/display_default_menu_item_setting", lc.FetchByID)
	group.POST("/display_default_menu_item_setting", lc.Create)
	group.POST("/display_default_menu_item_settings", lc.CreateMany)
	group.PUT("/display_default_menu_item_setting", lc.Update)
	group.DELETE("/display_default_menu_item_setting", lc.Delete)
}
