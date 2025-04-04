package route

import (
	"time"

	controller "earnforglance/server/api/controller/customers"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/customers"

	repository "earnforglance/server/repository/customers"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/customers"

	"github.com/gin-gonic/gin"
)

func CustomerRoleRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCustomerRoleRepository(db, domain.CollectionCustomerRole)
	lc := &controller.CustomerRoleController{
		CustomerRoleUsecase: usecase.NewCustomerRoleUsecase(ur, timeout),
		Env:                 env,
	}

	Group := group.Group("/api/v1/customers")

	Group.GET("/customer_roles", lc.Fetch)
	Group.GET("/customer_role", lc.FetchByID)
	Group.POST("/customer_role", lc.Create)
	Group.POST("/customer_roles", lc.CreateMany)
	Group.PUT("/customer_role", lc.Update)
	Group.DELETE("customer_role", lc.Delete)
}
