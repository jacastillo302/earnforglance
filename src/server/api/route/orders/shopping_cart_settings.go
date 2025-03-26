package route

import (
	"time"

	controller "earnforglance/server/api/controller/orders"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/orders"

	repository "earnforglance/server/repository/orders"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/orders"

	"github.com/gin-gonic/gin"
)

func ShoppingCartSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewShoppingCartSettingsRepository(db, domain.CollectionShoppingCartSettings)
	lc := &controller.ShoppingCartSettingsController{
		ShoppingCartSettingsUsecase: usecase.NewShoppingCartSettingsUsecase(ur, timeout),
		Env:                         env,
	}
	itemGroup := group.Group("/api/v1/orders")
	itemGroup.GET("/shopping_cart_settings", lc.Fetch)
	itemGroup.GET("/shopping_cart_setting", lc.FetchByID)
	itemGroup.POST("/shopping_cart_setting", lc.Create)
	itemGroup.POST("/shopping_cart_settings", lc.CreateMany)
	itemGroup.PUT("/shopping_cart_setting", lc.Update)
	itemGroup.DELETE("/shopping_cart_setting", lc.Delete)
}
