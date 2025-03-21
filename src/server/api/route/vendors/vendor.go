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

	group.GET("/vendors", lc.Fetch)
	group.GET("/vendor", lc.FetchByID)
	group.POST("/vendor", lc.Create)
	group.PUT("/vendor", lc.Update)
	group.DELETE("/vendor", lc.Delete)
}
