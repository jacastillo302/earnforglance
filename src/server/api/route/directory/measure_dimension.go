package route

import (
	"time"

	controller "earnforglance/server/api/controller/directory"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/directory"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/directory"
	usecase "earnforglance/server/usecase/directory"

	"github.com/gin-gonic/gin"
)

func MeasureDimensionRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewMeasureDimensionRepository(db, domain.CollectionMeasureDimension)
	lc := &controller.MeasureDimensionController{
		MeasureDimensionUsecase: usecase.NewMeasureDimensionUsecase(ur, timeout),
		Env:                     env,
	}

	itemGroup := group.Group("/api/v1/directory")
	itemGroup.GET("/measure_dimensions", lc.Fetch)
	itemGroup.GET("/measure_dimension", lc.FetchByID)
	itemGroup.POST("/measure_dimension", lc.Create)
	itemGroup.POST("/measure_dimensions", lc.CreateMany)
	itemGroup.PUT("/measure_dimension", lc.Update)
	itemGroup.DELETE("/measure_dimension", lc.Delete)
}
