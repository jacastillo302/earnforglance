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

	group.GET("/shipping_options", lc.Fetch)
	group.GET("/shipping_option", lc.FetchByID)
	group.POST("/shipping_option", lc.Create)
	group.POST("/shipping_options", lc.CreateMany)
	group.PUT("/shipping_option", lc.Update)
	group.DELETE("/shipping_option", lc.Delete)
}
