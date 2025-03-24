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

func ProductManufacturerRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductManufacturerRepository(db, domain.CollectionProductManufacturer)
	lc := &controller.ProductManufacturerController{
		ProductManufacturerUsecase: usecase.NewProductManufacturerUsecase(ur, timeout),
		Env:                        env,
	}

	group.GET("/product_manufacturers", lc.Fetch)
	group.GET("/product_manufacturer", lc.FetchByID)
	group.POST("/product_manufacturer", lc.Create)
	group.POST("/product_manufacturers", lc.CreateMany)
	group.PUT("/product_manufacturer", lc.Update)
	group.DELETE("/product_manufacturer", lc.Delete)
}
