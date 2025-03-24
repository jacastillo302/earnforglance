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

	group.GET("/warehouses", lc.Fetch)
	group.GET("/warehouse", lc.FetchByID)
	group.POST("/warehouse", lc.Create)
	group.POST("/warehouses", lc.CreateMany)
	group.PUT("/warehouse", lc.Update)
	group.DELETE("/warehouse", lc.Delete)
}
