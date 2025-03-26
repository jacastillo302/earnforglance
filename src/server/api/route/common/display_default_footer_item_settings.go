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

	itemGroup := group.Group("/api/v1/common")
	itemGroup.GET("/display_default_footer_item_settings", lc.Fetch)
	itemGroup.GET("/display_default_footer_item_setting", lc.FetchByID)
	itemGroup.POST("/display_default_footer_item_setting", lc.Create)
	itemGroup.POST("/display_default_footer_item_settings", lc.CreateMany)
	itemGroup.PUT("/display_default_footer_item_setting", lc.Update)
	itemGroup.DELETE("/display_default_footer_item_setting", lc.Delete)
}
