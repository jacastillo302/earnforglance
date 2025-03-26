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

func PickupPointRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPickupPointRepository(db, domain.CollectionPickupPoint)
	lc := &controller.PickupPointController{
		PickupPointUsecase: usecase.NewPickupPointUsecase(ur, timeout),
		Env:                env,
	}
	itemGroup := group.Group("/api/v1/shipping")
	itemGroup.GET("/pickup_points", lc.Fetch)
	itemGroup.GET("/pickup_point", lc.FetchByID)
	itemGroup.POST("/pickup_point", lc.Create)
	itemGroup.POST("/pickup_points", lc.CreateMany)
	itemGroup.PUT("/pickup_point", lc.Update)
	itemGroup.DELETE("/pickup_point", lc.Delete)
}
