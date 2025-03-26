package route

import (
	"time"

	controller "earnforglance/server/api/controller/common"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/common"

	repository "earnforglance/server/repository/common"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/common"

	"github.com/gin-gonic/gin"
)

func AddressSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewAddressSettingsRepository(db, domain.CollectionAddressSettings)
	lc := &controller.AddressSettingsController{
		AddressSettingsUsecase: usecase.NewAddressSettingsUsecase(ur, timeout),
		Env:                    env,
	}

	itemGroup := group.Group("/api/v1/common")
	itemGroup.GET("/address_settings", lc.Fetch)
	itemGroup.GET("/address_setting", lc.FetchByID)
	itemGroup.POST("/address_setting", lc.Create)
	itemGroup.POST("/address_settings", lc.CreateMany)
	itemGroup.PUT("/address_setting", lc.Update)
	itemGroup.DELETE("/address_setting", lc.Delete)
}
