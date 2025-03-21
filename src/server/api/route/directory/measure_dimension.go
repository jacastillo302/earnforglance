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

	group.GET("/measure_dimensions", lc.Fetch)
	group.GET("/measure_dimension", lc.FetchByID)
	group.POST("/measure_dimension", lc.Create)
	group.PUT("/measure_dimension", lc.Update)
	group.DELETE("/measure_dimension", lc.Delete)
}
