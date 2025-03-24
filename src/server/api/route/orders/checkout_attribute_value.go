package route

import (
	"time"

	controller "earnforglance/server/api/controller/orders"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/orders"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/orders"
	usecase "earnforglance/server/usecase/orders"

	"github.com/gin-gonic/gin"
)

func CheckoutAttributeValueRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCheckoutAttributeValueRepository(db, domain.CollectionCheckoutAttributeValue)
	lc := &controller.CheckoutAttributeValueController{
		CheckoutAttributeValueUsecase: usecase.NewCheckoutAttributeValueUsecase(ur, timeout),
		Env:                           env,
	}

	group.GET("/checkout_attribute_values", lc.Fetch)
	group.GET("/checkout_attribute_value", lc.FetchByID)
	group.POST("/checkout_attribute_value", lc.Create)
	group.POST("/checkout_attribute_values", lc.CreateMany)
	group.PUT("/checkout_attribute_value", lc.Update)
	group.DELETE("/checkout_attribute_value", lc.Delete)
}
