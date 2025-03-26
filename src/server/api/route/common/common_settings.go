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

	itemGroup := group.Group("/api/v1/common")
	itemGroup.GET("/common_settings", lc.Fetch)
	itemGroup.GET("/common_setting", lc.FetchByID)
	itemGroup.POST("/common_setting", lc.Create)
	itemGroup.POST("/common_settings", lc.CreateMany)
	itemGroup.PUT("/common_setting", lc.Update)
	itemGroup.DELETE("/common_setting", lc.Delete)
}
