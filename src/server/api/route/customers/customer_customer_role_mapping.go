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

	group.GET("/customer_customer_role_mappings", lc.Fetch)
	group.GET("/customer_customer_role_mapping", lc.FetchByID)
	group.POST("/customer_customer_role_mapping", lc.Create)
	group.PUT("/customer_customer_role_mapping", lc.Update)
	group.DELETE("/customer_customer_role_mapping", lc.Delete)
}
