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

func AdminAreaSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewAdminAreaSettingsRepository(db, domain.CollectionAdminAreaSettings)
	lc := &controller.AdminAreaSettingsController{
		AdminAreaSettingsUsecase: usecase.NewAdminAreaSettingsUsecase(ur, timeout),
		Env:                      env,
	}

	itemGroup := group.Group("/api/v1/common")
	itemGroup.GET("/admin_area_settings", lc.Fetch)
	itemGroup.GET("/admin_area_setting", lc.FetchByID)
	itemGroup.POST("/admin_area_setting", lc.Create)
	itemGroup.POST("/admin_area_settings", lc.CreateMany)
	itemGroup.PUT("/admin_area_setting", lc.Update)
	itemGroup.DELETE("/admin_area_setting", lc.Delete)
}
