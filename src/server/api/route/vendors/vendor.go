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

func VendorRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewVendorRepository(db, domain.CollectionVendor)
	lc := &controller.VendorController{
		VendorUsecase: usecase.NewVendorUsecase(ur, timeout),
		Env:           env,
	}
	itemGroup := group.Group("/api/v1/vendors")
	itemGroup.GET("/vendors", lc.Fetch)
	itemGroup.GET("/vendor", lc.FetchByID)
	itemGroup.POST("/vendor", lc.Create)
	itemGroup.POST("/vendors", lc.CreateMany)
	itemGroup.PUT("/vendor", lc.Update)
	itemGroup.DELETE("/vendor", lc.Delete)
}
