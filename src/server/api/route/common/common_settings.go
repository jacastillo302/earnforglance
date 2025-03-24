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

	group.GET("/common_settings", lc.Fetch)
	group.GET("/common_setting", lc.FetchByID)
	group.POST("/common_setting", lc.Create)
	group.POST("/common_settings", lc.CreateMany)
	group.PUT("/common_setting", lc.Update)
	group.DELETE("common_setting", lc.Delete)
}
