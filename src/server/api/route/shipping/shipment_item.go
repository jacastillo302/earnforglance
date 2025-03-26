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

func ShipmentItemRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewShipmentItemRepository(db, domain.CollectionShipmentItem)
	lc := &controller.ShipmentItemController{
		ShipmentItemUsecase: usecase.NewShipmentItemUsecase(ur, timeout),
		Env:                 env,
	}
	itemGroup := group.Group("/api/v1/shipping")
	itemGroup.GET("/shipment_items", lc.Fetch)
	itemGroup.GET("/shipment_item", lc.FetchByID)
	itemGroup.POST("/shipment_item", lc.Create)
	itemGroup.POST("/shipment_items", lc.CreateMany)
	itemGroup.PUT("/shipment_item", lc.Update)
	itemGroup.DELETE("/shipment_item", lc.Delete)
}
