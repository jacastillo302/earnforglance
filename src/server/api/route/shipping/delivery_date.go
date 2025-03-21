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

func DeliveryDateRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewDeliveryDateRepository(db, domain.CollectionDeliveryDate)
	lc := &controller.DeliveryDateController{
		DeliveryDateUsecase: usecase.NewDeliveryDateUsecase(ur, timeout),
		Env:                 env,
	}

	group.GET("/delivery_dates", lc.Fetch)
	group.GET("/delivery_date", lc.FetchByID)
	group.POST("/delivery_date", lc.Create)
	group.PUT("/delivery_date", lc.Update)
	group.DELETE("/delivery_date", lc.Delete)
}
