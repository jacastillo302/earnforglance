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

	group.GET("/customer_address_mappings", lc.Fetch)
	group.GET("/customer_address_mapping", lc.FetchByID)
	group.POST("/customer_address_mapping", lc.Create)
	group.POST("/customer_address_mappings", lc.CreateMany)
	group.PUT("/customer_address_mapping", lc.Update)
	group.DELETE("/customer_address_mapping", lc.Delete)
}
