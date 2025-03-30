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

func MultiFactorAuthenticationSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewMultiFactorAuthenticationSettingsRepository(db, domain.CollectionMultiFactorAuthenticationSettings)
	lc := &controller.MultiFactorAuthenticationSettingsController{
		MultiFactorAuthenticationSettingsUsecase: usecase.NewMultiFactorAuthenticationSettingsUsecase(ur, timeout),
		Env:                                      env,
	}

	Group := group.Group("/api/v1/customers")

	group.GET("/multiFactor_authentication_settings", lc.Fetch)
	group.GET("/multiFactor_authentication_setting", lc.FetchByID)
	group.POST("/multiFactor_authentication_setting", lc.Create)
	Group.POST("/multiFactor_authentication_settings", lc.CreateMany)
	group.PUT("/multiFactor_authentication_setting", lc.Update)
	group.DELETE("/multiFactor_authentication_setting", lc.Delete)
}
