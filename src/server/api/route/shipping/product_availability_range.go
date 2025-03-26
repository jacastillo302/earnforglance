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

func ProductAvailabilityRangeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductAvailabilityRangeRepository(db, domain.CollectionProductAvailabilityRange)
	lc := &controller.ProductAvailabilityRangeController{
		ProductAvailabilityRangeUsecase: usecase.NewProductAvailabilityRangeUsecase(ur, timeout),
		Env:                             env,
	}
	itemGroup := group.Group("/api/v1/shipping")
	itemGroup.GET("/product_availability_ranges", lc.Fetch)
	itemGroup.GET("/product_availability_range", lc.FetchByID)
	itemGroup.POST("/product_availability_range", lc.Create)
	itemGroup.POST("/product_availability_ranges", lc.CreateMany)
	itemGroup.PUT("/product_availability_range", lc.Update)
	itemGroup.DELETE("/product_availability_range", lc.Delete)
}
