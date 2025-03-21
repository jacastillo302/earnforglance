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

	group.GET("/cross_sell_products", lc.Fetch)
	group.GET("/cross_sell_product", lc.FetchByID)
	group.POST("/cross_sell_product", lc.Create)
	group.PUT("/cross_sell_product", lc.Update)
	group.DELETE("/cross_sell_product", lc.Delete)
}
