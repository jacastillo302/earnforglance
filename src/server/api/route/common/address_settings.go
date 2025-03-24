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

func AddressSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewAddressSettingsRepository(db, domain.CollectionAddressSettings)
	lc := &controller.AddressSettingsController{
		AddressSettingsUsecase: usecase.NewAddressSettingsUsecase(ur, timeout),
		Env:                    env,
	}

	group.GET("/address_settings", lc.Fetch)
	group.GET("/address_setting", lc.FetchByID)
	group.POST("/address_setting", lc.Create)
	group.POST("/address_settings", lc.CreateMany)
	group.PUT("/address_setting", lc.Update)
	group.DELETE("address_setting", lc.Delete)
}
