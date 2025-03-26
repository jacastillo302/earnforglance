package route

import (
	"time"

	controller "earnforglance/server/api/controller/catalog"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/catalog"

	repository "earnforglance/server/repository/catalog"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/catalog"

	"github.com/gin-gonic/gin"
)

func StockQuantityChangeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewStockQuantityChangeRepository(db, domain.CollectionStockQuantityChange)
	lc := &controller.StockQuantityChangeController{
		StockQuantityChangeUsecase: usecase.NewStockQuantityChangeUsecase(ur, timeout),
		Env:                        env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/stock_quantity_changes", lc.Fetch)
	itemGroup.GET("/stock_quantity_change", lc.FetchByID)
	itemGroup.POST("/stock_quantity_change", lc.Create)
	itemGroup.POST("/stock_quantity_changes", lc.CreateMany)
	itemGroup.PUT("/stock_quantity_change", lc.Update)
	itemGroup.DELETE("/stock_quantity_change", lc.Delete)
}
