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

func ProxySettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProxySettingsRepository(db, domain.CollectionProxySettings)
	lc := &controller.ProxySettingsController{
		ProxySettingsUsecase: usecase.NewProxySettingsUsecase(ur, timeout),
		Env:                  env,
	}

	group.GET("/proxy_settings", lc.Fetch)
	group.GET("/proxy_setting", lc.FetchByID)
	group.POST("/proxy_setting", lc.Create)
	group.PUT("/proxy_setting", lc.Update)
	group.DELETE("proxy_setting", lc.Delete)
}
