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

func ShoppingCartItemRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewShoppingCartItemRepository(db, domain.CollectionShoppingCartItem)
	lc := &controller.ShoppingCartItemController{
		ShoppingCartItemUsecase: usecase.NewShoppingCartItemUsecase(ur, timeout),
		Env:                     env,
	}
	itemGroup := group.Group("/api/v1/orders")
	itemGroup.GET("/shoppingcartitems", lc.Fetch)
	itemGroup.GET("/shoppingcartitem", lc.FetchByID)
	itemGroup.POST("/shoppingcartitem", lc.Create)
	itemGroup.POST("/shoppingcartitems", lc.CreateMany)
	itemGroup.PUT("/shoppingcartitem", lc.Update)
	itemGroup.DELETE("/shoppingcartitem", lc.Delete)
}
