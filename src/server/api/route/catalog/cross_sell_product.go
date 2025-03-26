package route

import (
	"time"

	controller "earnforglance/server/api/controller/catalog"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/catalog"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/catalog"
	usecase "earnforglance/server/usecase/catalog"

	"github.com/gin-gonic/gin"
)

func CrossSellProductRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCrossSellProductRepository(db, domain.CollectionCrossSellProduct)
	lc := &controller.CrossSellProductController{
		CrossSellProductUsecase: usecase.NewCrossSellProductUsecase(ur, timeout),
		Env:                     env,
	}

	itemGroup := group.Group("/api/v1/catalog")

	itemGroup.Group("/" + domain.CollectionCrossSellProduct)
	itemGroup.GET("/cross_sell_products", lc.Fetch)
	itemGroup.GET("/cross_sell_product", lc.FetchByID)
	itemGroup.POST("/cross_sell_product", lc.Create)
	itemGroup.POST("/cross_sell_products", lc.CreateMany)
	itemGroup.PUT("/cross_sell_product", lc.Update)
	itemGroup.DELETE("/cross_sell_product", lc.Delete)
}
