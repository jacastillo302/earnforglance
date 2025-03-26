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

func VendorAttributeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewVendorAttributeRepository(db, domain.CollectionVendorAttribute)
	lc := &controller.VendorAttributeController{
		VendorAttributeUsecase: usecase.NewVendorAttributeUsecase(ur, timeout),
		Env:                    env,
	}
	itemGroup := group.Group("/api/v1/vendors")
	itemGroup.GET("/vendor_attributes", lc.Fetch)
	itemGroup.GET("/vendor_attribute", lc.FetchByID)
	itemGroup.POST("/vendor_attribute", lc.Create)
	itemGroup.POST("/vendor_attributes", lc.CreateMany)
	itemGroup.PUT("/vendor_attribute", lc.Update)
	itemGroup.DELETE("/vendor_attribute", lc.Delete)
}
