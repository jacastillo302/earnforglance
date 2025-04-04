package route

import (
	"time"

	controller "earnforglance/server/api/controller/customers"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/customers"

	repository "earnforglance/server/repository/customers"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/customers"

	"github.com/gin-gonic/gin"
)

func ExternalAuthenticationSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewExternalAuthenticationSettingsRepository(db, domain.CollectionExternalAuthenticationSettings)
	lc := &controller.ExternalAuthenticationSettingsController{
		ExternalAuthenticationSettingsUsecase: usecase.NewExternalAuthenticationSettingsUsecase(ur, timeout),
		Env:                                   env,
	}

	Group := group.Group("/api/v1/customers")

	Group.GET("/external_authentication_settings", lc.Fetch)
	Group.GET("/external_authentication_setting", lc.FetchByID)
	Group.POST("/external_authentication_setting", lc.Create)
	Group.POST("/external_authentication_settings", lc.CreateMany)
	Group.PUT("/external_authentication_setting", lc.Update)
	Group.DELETE("external_authentication_setting", lc.Delete)
}
