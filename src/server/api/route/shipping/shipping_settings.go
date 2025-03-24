package route

import (
	"time"

	controller "earnforglance/server/api/controller/shipping"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/shipping"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/shipping"
	usecase "earnforglance/server/usecase/shipping"

	"github.com/gin-gonic/gin"
)

func ShippingSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewShippingSettingsRepository(db, domain.CollectionShippingSettings)
	lc := &controller.ShippingSettingsController{
		ShippingSettingsUsecase: usecase.NewShippingSettingsUsecase(ur, timeout),
		Env:                     env,
	}

	group.GET("/shipping_settings", lc.Fetch)
	group.GET("/shipping_setting", lc.FetchByID)
	group.POST("/shipping_setting", lc.Create)
	group.POST("/shipping_settings", lc.CreateMany)
	group.PUT("/shipping_setting", lc.Update)
	group.DELETE("/shipping_setting", lc.Delete)
}
