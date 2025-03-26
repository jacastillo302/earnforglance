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

func ProductManufacturerRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductManufacturerRepository(db, domain.CollectionProductManufacturer)
	lc := &controller.ProductManufacturerController{
		ProductManufacturerUsecase: usecase.NewProductManufacturerUsecase(ur, timeout),
		Env:                        env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/product_manufacturers", lc.Fetch)
	itemGroup.GET("/product_manufacturer", lc.FetchByID)
	itemGroup.POST("/product_manufacturer", lc.Create)
	itemGroup.POST("/product_manufacturers", lc.CreateMany)
	itemGroup.PUT("/product_manufacturer", lc.Update)
	itemGroup.DELETE("/product_manufacturer", lc.Delete)
}
