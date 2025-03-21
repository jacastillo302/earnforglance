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

func ShoppingCartSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewShoppingCartSettingsRepository(db, domain.CollectionShoppingCartSettings)
	lc := &controller.ShoppingCartSettingsController{
		ShoppingCartSettingsUsecase: usecase.NewShoppingCartSettingsUsecase(ur, timeout),
		Env:                         env,
	}

	group.GET("/shopping_cart_settingss", lc.Fetch)
	group.GET("/shopping_cart_settings", lc.FetchByID)
	group.POST("/shopping_cart_settings", lc.Create)
	group.PUT("/shopping_cart_settings", lc.Update)
	group.DELETE("/shopping_cart_settings", lc.Delete)
}
