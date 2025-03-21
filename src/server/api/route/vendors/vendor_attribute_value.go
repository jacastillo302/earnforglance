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

func VendorAttributeValueRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewVendorAttributeValueRepository(db, domain.CollectionVendorAttributeValue)
	lc := &controller.VendorAttributeValueController{
		VendorAttributeValueUsecase: usecase.NewVendorAttributeValueUsecase(ur, timeout),
		Env:                         env,
	}

	group.GET("/vendor_attribute_values", lc.Fetch)
	group.GET("/vendor_attribute_value", lc.FetchByID)
	group.POST("/vendor_attribute_value", lc.Create)
	group.PUT("/vendor_attribute_value", lc.Update)
	group.DELETE("/vendor_attribute_value", lc.Delete)
}
