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
	itemGroup := group.Group("/api/v1/shipping")
	itemGroup.GET("/shipping_settings", lc.Fetch)
	itemGroup.GET("/shipping_setting", lc.FetchByID)
	itemGroup.POST("/shipping_setting", lc.Create)
	itemGroup.POST("/shipping_settings", lc.CreateMany)
	itemGroup.PUT("/shipping_setting", lc.Update)
	itemGroup.DELETE("/shipping_setting", lc.Delete)
}
