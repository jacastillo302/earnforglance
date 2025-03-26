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

	itemGroup := group.Group("/api/v1/directory")
	itemGroup.GET("/measure_weights", lc.Fetch)
	itemGroup.GET("/measure_weight", lc.FetchByID)
	itemGroup.POST("/measure_weight", lc.Create)
	itemGroup.POST("/measure_weights", lc.CreateMany)
	itemGroup.PUT("/measure_weight", lc.Update)
	itemGroup.DELETE("/measure_weight", lc.Delete)
}
