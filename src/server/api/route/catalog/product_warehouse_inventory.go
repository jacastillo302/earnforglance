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

func ProductWarehouseInventoryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductWarehouseInventoryRepository(db, domain.CollectionProductWarehouseInventory)
	lc := &controller.ProductWarehouseInventoryController{
		ProductWarehouseInventoryUsecase: usecase.NewProductWarehouseInventoryUsecase(ur, timeout),
		Env:                              env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/product_warehouse_inventories", lc.Fetch)
	itemGroup.GET("/product_warehouse_inventory", lc.FetchByID)
	itemGroup.POST("/product_warehouse_inventory", lc.Create)
	itemGroup.POST("/product_warehouse_inventories", lc.CreateMany)
	itemGroup.PUT("/product_warehouse_inventory", lc.Update)
	itemGroup.DELETE("/product_warehouse_inventory", lc.Delete)
}
