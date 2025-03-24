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

	group.GET("/checkout_attributes", lc.Fetch)
	group.GET("/checkout_attribute", lc.FetchByID)
	group.POST("/checkout_attribute", lc.Create)
	group.POST("/checkout_attributes", lc.CreateMany)
	group.PUT("/checkout_attribute", lc.Update)
	group.DELETE("/checkout_attribute", lc.Delete)
}
