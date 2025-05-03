package route

import (
	"time"

	controller "earnforglance/server/api/controller/attributes"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/attributes"

	repository "earnforglance/server/repository/attributes"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/attributes"

	"github.com/gin-gonic/gin"
)

func CustomerAttributeValueRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCustomerAttributeValueRepository(db, domain.CollectionCustomerAttributeValue)
	lc := &controller.CustomerAttributeValueController{
		CustomerAttributeValueUsecase: usecase.NewCustomerAttributeValueUsecase(ur, timeout),
		Env:                           env,
	}

	itemGroup := group.Group("/api/v1/customers")
	itemGroup.GET("/customer_attribute_values", lc.Fetch)
	itemGroup.GET("/customer_attribute_value", lc.FetchByID)
	itemGroup.POST("/customer_attribute_value", lc.Create)
	itemGroup.POST("/customer_attribute_values", lc.CreateMany)
	itemGroup.PUT("/customer_attribute_value", lc.Update)
	itemGroup.DELETE("/customer_attribute_value", lc.Delete)
}
