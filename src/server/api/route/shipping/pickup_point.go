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

	group.GET("/pickup_points", lc.Fetch)
	group.GET("/pickup_point", lc.FetchByID)
	group.POST("/pickup_point", lc.Create)
	group.PUT("/pickup_point", lc.Update)
	group.DELETE("/pickup_point", lc.Delete)
}
