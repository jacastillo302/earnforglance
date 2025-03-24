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

	group.GET("/product_availability_ranges", lc.Fetch)
	group.GET("/product_availability_range", lc.FetchByID)
	group.POST("/product_availability_range", lc.Create)
	group.POST("/product_availability_ranges", lc.CreateMany)
	group.PUT("/product_availability_range", lc.Update)
	group.DELETE("/product_availability_range", lc.Delete)
}
