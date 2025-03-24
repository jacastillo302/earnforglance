package route

import (
	"time"

	controller "earnforglance/server/api/controller/vendors"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/vendors"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/vendors"
	usecase "earnforglance/server/usecase/vendors"

	"github.com/gin-gonic/gin"
)

func VendorSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewVendorSettingsRepository(db, domain.CollectionVendorSettings)
	lc := &controller.VendorSettingsController{
		VendorSettingsUsecase: usecase.NewVendorSettingsUsecase(ur, timeout),
		Env:                   env,
	}

	group.GET("/vendor_settings", lc.Fetch)
	group.GET("/vendor_setting", lc.FetchByID)
	group.POST("/vendor_setting", lc.Create)
	group.POST("/vendor_settings", lc.CreateMany)
	group.PUT("/vendor_setting", lc.Update)
	group.DELETE("vendor_setting", lc.Delete)
}
