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

func DeliveryDateRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewDeliveryDateRepository(db, domain.CollectionDeliveryDate)
	lc := &controller.DeliveryDateController{
		DeliveryDateUsecase: usecase.NewDeliveryDateUsecase(ur, timeout),
		Env:                 env,
	}
	itemGroup := group.Group("/api/v1/shipping")
	itemGroup.GET("/delivery_dates", lc.Fetch)
	itemGroup.GET("/delivery_date", lc.FetchByID)
	itemGroup.POST("/delivery_date", lc.Create)
	itemGroup.POST("/delivery_dates", lc.CreateMany)
	itemGroup.PUT("/delivery_date", lc.Update)
	itemGroup.DELETE("/delivery_date", lc.Delete)
}
