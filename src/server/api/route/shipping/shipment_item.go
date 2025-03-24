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

func ShipmentItemRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewShipmentItemRepository(db, domain.CollectionShipmentItem)
	lc := &controller.ShipmentItemController{
		ShipmentItemUsecase: usecase.NewShipmentItemUsecase(ur, timeout),
		Env:                 env,
	}

	group.GET("/shipment_items", lc.Fetch)
	group.GET("/shipment_item", lc.FetchByID)
	group.POST("/shipment_item", lc.Create)
	group.POST("/shipment_items", lc.CreateMany)
	group.PUT("/shipment_item", lc.Update)
	group.DELETE("/shipment_item", lc.Delete)
}
