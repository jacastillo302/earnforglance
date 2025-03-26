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
	itemGroup := group.Group("/api/v1/gdpr")
	itemGroup.GET("/gdpr_settings", lc.Fetch)
	itemGroup.GET("/gdpr_setting", lc.FetchByID)
	itemGroup.POST("/gdpr_setting", lc.Create)
	itemGroup.POST("/gdpr_settings", lc.CreateMany)
	itemGroup.PUT("/gdpr_setting", lc.Update)
	itemGroup.DELETE("/gdpr_setting", lc.Delete)
}
