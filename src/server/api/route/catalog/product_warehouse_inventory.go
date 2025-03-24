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

	group.GET("/product_warehouse_inventories", lc.Fetch)
	group.GET("/product_warehouse_inventory", lc.FetchByID)
	group.POST("/product_warehouse_inventory", lc.Create)
	group.POST("/product_warehouse_inventories", lc.CreateMany)
	group.PUT("/product_warehouse_inventory", lc.Update)
	group.DELETE("/product_warehouse_inventory", lc.Delete)
}
