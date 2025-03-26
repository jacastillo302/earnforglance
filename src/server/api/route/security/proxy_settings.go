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
	itemGroup := group.Group("/api/v1/security")
	itemGroup.GET("/proxy_settings", lc.Fetch)
	itemGroup.GET("/proxy_setting", lc.FetchByID)
	itemGroup.POST("/proxy_setting", lc.Create)
	itemGroup.POST("/proxy_settings", lc.CreateMany)
	itemGroup.PUT("/proxy_setting", lc.Update)
	itemGroup.DELETE("/proxy_setting", lc.Delete)
}
