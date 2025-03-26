package route

import (
	"time"

	controller "earnforglance/server/api/controller/customers"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/customers"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/customers"
	usecase "earnforglance/server/usecase/customers"

	"github.com/gin-gonic/gin"
)

func CustomerAddressMappingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCustomerAddressMappingRepository(db, domain.CollectionCustomerAddressMapping)
	lc := &controller.CustomerAddressMappingController{
		CustomerAddressMappingUsecase: usecase.NewCustomerAddressMappingUsecase(ur, timeout),
		Env:                           env,
	}

	itemGroup := group.Group("/api/v1/customers")
	itemGroup.GET("/customer_address_mappings", lc.Fetch)
	itemGroup.GET("/customer_address_mapping", lc.FetchByID)
	itemGroup.POST("/customer_address_mapping", lc.Create)
	itemGroup.POST("/customer_address_mappings", lc.CreateMany)
	itemGroup.PUT("/customer_address_mapping", lc.Update)
	itemGroup.DELETE("/customer_address_mapping", lc.Delete)
}
