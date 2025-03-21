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

func DisplayDefaultFooterItemSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewDisplayDefaultFooterItemSettingsRepository(db, domain.CollectionDisplayDefaultFooterItemSettings)
	lc := &controller.DisplayDefaultFooterItemSettingsController{
		DisplayDefaultFooterItemSettingsUsecase: usecase.NewDisplayDefaultFooterItemSettingsUsecase(ur, timeout),
		Env:                                     env,
	}

	group.GET("/display_default_footer_item_settings", lc.Fetch)
	group.GET("/display_default_footer_item_setting", lc.FetchByID)
	group.POST("/display_default_footer_item_setting", lc.Create)
	group.PUT("/display_default_footer_item_setting", lc.Update)
	group.DELETE("/display_default_footer_item_setting", lc.Delete)
}
