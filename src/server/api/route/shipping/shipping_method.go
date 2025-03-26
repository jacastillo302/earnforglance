package route

import (
	"time"

	controller "earnforglance/server/api/controller/shipping"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/shipping"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/shipping"
	usecase "earnforglance/server/usecase/shipping"

	"github.com/gin-gonic/gin"
)

func ShippingMethodRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewShippingMethodRepository(db, domain.CollectionShippingMethod)
	lc := &controller.ShippingMethodController{
		ShippingMethodUsecase: usecase.NewShippingMethodUsecase(ur, timeout),
		Env:                   env,
	}
	itemGroup := group.Group("/api/v1/shipping")
	itemGroup.GET("/shipping_methods", lc.Fetch)
	itemGroup.GET("/shipping_method", lc.FetchByID)
	itemGroup.POST("/shipping_method", lc.Create)
	itemGroup.POST("/shipping_methods", lc.CreateMany)
	itemGroup.PUT("/shipping_method", lc.Update)
	itemGroup.DELETE("/shipping_method", lc.Delete)
}
