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

func CommonSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCommonSettingsRepository(db, domain.CollectionCommonSettings)
	lc := &controller.CommonSettingsController{
		CommonSettingsUsecase: usecase.NewCommonSettingsUsecase(ur, timeout),
		Env:                   env,
	}

	group.GET("/common_settingss", lc.Fetch)
	group.GET("/common_settings", lc.FetchByID)
	group.POST("/common_settings", lc.Create)
	group.PUT("/common_settings", lc.Update)
	group.DELETE("common_settings", lc.Delete)
}
