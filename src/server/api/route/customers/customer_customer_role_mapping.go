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

func CustomerCustomerRoleMappingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCustomerCustomerRoleMappingRepository(db, domain.CollectionCustomerCustomerRoleMapping)
	lc := &controller.CustomerCustomerRoleMappingController{
		CustomerCustomerRoleMappingUsecase: usecase.NewCustomerCustomerRoleMappingUsecase(ur, timeout),
		Env:                                env,
	}

	itemGroup := group.Group("/api/v1/customers")
	itemGroup.GET("/customer_customer_role_mappings", lc.Fetch)
	itemGroup.GET("/customer_customer_role_mapping", lc.FetchByID)
	itemGroup.POST("/customer_customer_role_mapping", lc.Create)
	itemGroup.POST("/customer_customer_role_mappings", lc.CreateMany)
	itemGroup.PUT("/customer_customer_role_mapping", lc.Update)
	itemGroup.DELETE("/customer_customer_role_mapping", lc.Delete)
}
