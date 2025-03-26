package route

import (
	"time"

	controller "earnforglance/server/api/controller/shipping"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/shipping"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/shipping"
	usecase "earnforglance/server/usecase/shipping"

	"github.com/gin-gonic/gin"
)

func WarehouseRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewWarehouseRepository(db, domain.CollectionWarehouse)
	lc := &controller.WarehouseController{
		WarehouseUsecase: usecase.NewWarehouseUsecase(ur, timeout),
		Env:              env,
	}
	itemGroup := group.Group("/api/v1/shipping")
	itemGroup.GET("/warehouses", lc.Fetch)
	itemGroup.GET("/warehouse", lc.FetchByID)
	itemGroup.POST("/warehouse", lc.Create)
	itemGroup.POST("/warehouses", lc.CreateMany)
	itemGroup.PUT("/warehouse", lc.Update)
	itemGroup.DELETE("/warehouse", lc.Delete)
}
