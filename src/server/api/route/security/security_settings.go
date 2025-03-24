package route

import (
	"time"

	controller "earnforglance/server/api/controller/security"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/security"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/security"
	usecase "earnforglance/server/usecase/security"

	"github.com/gin-gonic/gin"
)

func SecuritySettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSecuritySettingsRepository(db, domain.CollectionSecuritySettings)
	lc := &controller.SecuritySettingsController{
		SecuritySettingsUsecase: usecase.NewSecuritySettingsUsecase(ur, timeout),
		Env:                     env,
	}

	group.GET("/security_settings", lc.Fetch)
	group.GET("/security_setting", lc.FetchByID)
	group.POST("/security_setting", lc.Create)
	group.POST("/security_settings", lc.CreateMany)
	group.PUT("/security_setting", lc.Update)
	group.DELETE("security_setting", lc.Delete)
}
