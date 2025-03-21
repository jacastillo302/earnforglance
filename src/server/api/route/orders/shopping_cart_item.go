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

func ShoppingCartItemRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewShoppingCartItemRepository(db, domain.CollectionShoppingCartItem)
	lc := &controller.ShoppingCartItemController{
		ShoppingCartItemUsecase: usecase.NewShoppingCartItemUsecase(ur, timeout),
		Env:                     env,
	}

	group.GET("/shoppingcartitems", lc.Fetch)
	group.GET("/shoppingcartitem", lc.FetchByID)
	group.POST("/shoppingcartitem", lc.Create)
	group.PUT("/shoppingcartitem", lc.Update)
	group.DELETE("/shoppingcartitem", lc.Delete)
}
