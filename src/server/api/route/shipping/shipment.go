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

func ShipmentRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewShipmentRepository(db, domain.CollectionShipment)
	lc := &controller.ShipmentController{
		ShipmentUsecase: usecase.NewShipmentUsecase(ur, timeout),
		Env:             env,
	}

	group.GET("/shipments", lc.Fetch)
	group.GET("/shipment", lc.FetchByID)
	group.POST("/shipment", lc.Create)
	group.PUT("/shipment", lc.Update)
	group.DELETE("/shipment", lc.Delete)
}
