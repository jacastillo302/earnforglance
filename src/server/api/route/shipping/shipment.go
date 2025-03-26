package route

import (
	"time"

	controller "earnforglance/server/api/controller/shipping"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/shipping"

	repository "earnforglance/server/repository/shipping"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/shipping"

	"github.com/gin-gonic/gin"
)

func ShipmentRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewShipmentRepository(db, domain.CollectionShipment)
	lc := &controller.ShipmentController{
		ShipmentUsecase: usecase.NewShipmentUsecase(ur, timeout),
		Env:             env,
	}
	itemGroup := group.Group("/api/v1/shipping")
	itemGroup.GET("/shipments", lc.Fetch)
	itemGroup.GET("/shipment", lc.FetchByID)
	itemGroup.POST("/shipment", lc.Create)
	itemGroup.POST("/shipments", lc.CreateMany)
	itemGroup.PUT("/shipment", lc.Update)
	itemGroup.DELETE("/shipment", lc.Delete)
}
