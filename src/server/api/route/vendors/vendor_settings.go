package route

import (
	"time"

	controller "earnforglance/server/api/controller/vendors"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/vendors"

	repository "earnforglance/server/repository/vendors"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/vendors"

	"github.com/gin-gonic/gin"
)

func VendorSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewVendorSettingsRepository(db, domain.CollectionVendorSettings)
	lc := &controller.VendorSettingsController{
		VendorSettingsUsecase: usecase.NewVendorSettingsUsecase(ur, timeout),
		Env:                   env,
	}
	itemGroup := group.Group("/api/v1/vendors")
	itemGroup.GET("/vendor_settings", lc.Fetch)
	itemGroup.GET("/vendor_setting", lc.FetchByID)
	itemGroup.POST("/vendor_setting", lc.Create)
	itemGroup.POST("/vendor_settings", lc.CreateMany)
	itemGroup.PUT("/vendor_setting", lc.Update)
	itemGroup.DELETE("vendor_setting", lc.Delete)
}
