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

func StockQuantityChangeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewStockQuantityChangeRepository(db, domain.CollectionStockQuantityChange)
	lc := &controller.StockQuantityChangeController{
		StockQuantityChangeUsecase: usecase.NewStockQuantityChangeUsecase(ur, timeout),
		Env:                        env,
	}

	group.GET("/stock_quantity_changes", lc.Fetch)
	group.GET("/stock_quantity_change", lc.FetchByID)
	group.POST("/stock_quantity_change", lc.Create)
	group.POST("/stock_quantity_changes", lc.CreateMany)
	group.PUT("/stock_quantity_change", lc.Update)
	group.DELETE("/stock_quantity_change", lc.Delete)
}
