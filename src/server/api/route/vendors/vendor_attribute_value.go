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

func VendorAttributeValueRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewVendorAttributeValueRepository(db, domain.CollectionVendorAttributeValue)
	lc := &controller.VendorAttributeValueController{
		VendorAttributeValueUsecase: usecase.NewVendorAttributeValueUsecase(ur, timeout),
		Env:                         env,
	}
	itemGroup := group.Group("/api/v1/vendors")
	itemGroup.GET("/vendor_attribute_values", lc.Fetch)
	itemGroup.GET("/vendor_attribute_value", lc.FetchByID)
	itemGroup.POST("/vendor_attribute_value", lc.Create)
	itemGroup.POST("/vendor_attribute_values", lc.CreateMany)
	itemGroup.PUT("/vendor_attribute_value", lc.Update)
	itemGroup.DELETE("/vendor_attribute_value", lc.Delete)
}
