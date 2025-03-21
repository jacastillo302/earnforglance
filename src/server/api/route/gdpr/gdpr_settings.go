package route

import (
	"time"

	controller "earnforglance/server/api/controller/gdpr"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/gdpr"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/gdpr"
	usecase "earnforglance/server/usecase/gdpr"

	"github.com/gin-gonic/gin"
)

func GdprSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewGdprSettingsRepository(db, domain.CollectionGdprSettings)
	lc := &controller.GdprSettingsController{
		GdprSettingsUsecase: usecase.NewGdprSettingsUsecase(ur, timeout),
		Env:                 env,
	}

	group.GET("/gdpr_settingss", lc.Fetch)
	group.GET("/gdpr_settings", lc.FetchByID)
	group.POST("/gdpr_settings", lc.Create)
	group.PUT("/gdpr_settings", lc.Update)
	group.DELETE("/gdpr_settings", lc.Delete)
}
