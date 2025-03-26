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

func ShippingOptionRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewShippingOptionRepository(db, domain.CollectionShippingOption)
	lc := &controller.ShippingOptionController{
		ShippingOptionUsecase: usecase.NewShippingOptionUsecase(ur, timeout),
		Env:                   env,
	}
	itemGroup := group.Group("/api/v1/shipping")
	itemGroup.GET("/shipping_options", lc.Fetch)
	itemGroup.GET("/shipping_option", lc.FetchByID)
	itemGroup.POST("/shipping_option", lc.Create)
	itemGroup.POST("/shipping_options", lc.CreateMany)
	itemGroup.PUT("/shipping_option", lc.Update)
	itemGroup.DELETE("/shipping_option", lc.Delete)
}
