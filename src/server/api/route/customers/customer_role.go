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

	group.GET("/customer_roles", lc.Fetch)
	group.GET("/customer_role", lc.FetchByID)
	group.POST("/customer_role", lc.Create)
	Group.POST("/customer_roles", lc.CreateMany)
	group.PUT("/customer_role", lc.Update)
	group.DELETE("customer_role", lc.Delete)
}
