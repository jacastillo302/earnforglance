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

func MeasureWeightRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewMeasureWeightRepository(db, domain.CollectionMeasureWeight)
	lc := &controller.MeasureWeightController{
		MeasureWeightUsecase: usecase.NewMeasureWeightUsecase(ur, timeout),
		Env:                  env,
	}

	group.GET("/measure_weights", lc.Fetch)
	group.GET("/measure_weight", lc.FetchByID)
	group.POST("/measure_weight", lc.Create)
	group.PUT("/measure_weight", lc.Update)
	group.DELETE("/measure_weight", lc.Delete)
}
