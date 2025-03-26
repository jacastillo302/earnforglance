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

func ShippingMethodCountryMappingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewShippingMethodCountryMappingRepository(db, domain.CollectionShippingMethodCountryMapping)
	lc := &controller.ShippingMethodCountryMappingController{
		ShippingMethodCountryMappingUsecase: usecase.NewShippingMethodCountryMappingUsecase(ur, timeout),
		Env:                                 env,
	}
	itemGroup := group.Group("/api/v1/shipping")
	itemGroup.GET("/shipping_method_country_mappings", lc.Fetch)
	itemGroup.GET("/shipping_method_country_mapping", lc.FetchByID)
	itemGroup.POST("/shipping_method_country_mapping", lc.Create)
	itemGroup.POST("/shipping_method_country_mappings", lc.CreateMany)
	itemGroup.PUT("/shipping_method_country_mapping", lc.Update)
	itemGroup.DELETE("/shipping_method_country_mapping", lc.Delete)
}
