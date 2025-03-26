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

func CheckoutAttributeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCheckoutAttributeRepository(db, domain.CollectionCheckoutAttribute)
	lc := &controller.CheckoutAttributeController{
		CheckoutAttributeUsecase: usecase.NewCheckoutAttributeUsecase(ur, timeout),
		Env:                      env,
	}
	itemGroup := group.Group("/api/v1/orders")
	itemGroup.GET("/checkout_attributes", lc.Fetch)
	itemGroup.GET("/checkout_attribute", lc.FetchByID)
	itemGroup.POST("/checkout_attribute", lc.Create)
	itemGroup.POST("/checkout_attributes", lc.CreateMany)
	itemGroup.PUT("/checkout_attribute", lc.Update)
	itemGroup.DELETE("/checkout_attribute", lc.Delete)
}
