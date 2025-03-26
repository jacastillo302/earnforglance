package route

import (
	"time"

	controller "earnforglance/server/api/controller/security"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/security"

	repository "earnforglance/server/repository/security"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/security"

	"github.com/gin-gonic/gin"
)

func SecuritySettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSecuritySettingsRepository(db, domain.CollectionSecuritySettings)
	lc := &controller.SecuritySettingsController{
		SecuritySettingsUsecase: usecase.NewSecuritySettingsUsecase(ur, timeout),
		Env:                     env,
	}
	itemGroup := group.Group("/api/v1/security")
	itemGroup.GET("/security_settings", lc.Fetch)
	itemGroup.GET("/security_setting", lc.FetchByID)
	itemGroup.POST("/security_setting", lc.Create)
	itemGroup.POST("/security_settings", lc.CreateMany)
	itemGroup.PUT("/security_setting", lc.Update)
	itemGroup.DELETE("/security_setting", lc.Delete)
}
