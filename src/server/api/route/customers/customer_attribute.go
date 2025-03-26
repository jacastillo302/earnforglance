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

	itemGroup := group.Group("/api/v1/customers")
	itemGroup.GET("/customer_attributes", lc.Fetch)
	itemGroup.GET("/customer_attribute", lc.FetchByID)
	itemGroup.POST("/customer_attribute", lc.Create)
	itemGroup.POST("/customer_attributes", lc.CreateMany)
	itemGroup.PUT("/customer_attribute", lc.Update)
	itemGroup.DELETE("/customer_attribute", lc.Delete)
}
