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

func CustomerAttributeValueRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCustomerAttributeValueRepository(db, domain.CollectionCustomerAttributeValue)
	lc := &controller.CustomerAttributeValueController{
		CustomerAttributeValueUsecase: usecase.NewCustomerAttributeValueUsecase(ur, timeout),
		Env:                           env,
	}

	group.GET("/customer_attribute_values", lc.Fetch)
	group.GET("/customer_attribute_value", lc.FetchByID)
	group.POST("/customer_attribute_value", lc.Create)
	group.POST("/customer_attribute_values", lc.CreateMany)
	group.PUT("/customer_attribute_value", lc.Update)
	group.DELETE("/customer_attribute_value", lc.Delete)
}
