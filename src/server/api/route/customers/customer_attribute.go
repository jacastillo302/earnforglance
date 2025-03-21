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

func CustomerAttributeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCustomerAttributeRepository(db, domain.CollectionCustomerAttribute)
	lc := &controller.CustomerAttributeController{
		CustomerAttributeUsecase: usecase.NewCustomerAttributeUsecase(ur, timeout),
		Env:                      env,
	}

	group.GET("/customer_attributes", lc.Fetch)
	group.GET("/customer_attribute", lc.FetchByID)
	group.POST("/customer_attribute", lc.Create)
	group.PUT("/customer_attribute", lc.Update)
	group.DELETE("/customer_attribute", lc.Delete)
}
