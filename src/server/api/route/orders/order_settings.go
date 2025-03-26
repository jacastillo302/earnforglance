package route

import (
	"time"

	controller "earnforglance/server/api/controller/orders"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/orders"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/orders"
	usecase "earnforglance/server/usecase/orders"

	"github.com/gin-gonic/gin"
)

func OrderSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewOrderSettingsRepository(db, domain.CollectionOrderSettings)
	lc := &controller.OrderSettingsController{
		OrderSettingsUsecase: usecase.NewOrderSettingsUsecase(ur, timeout),
		Env:                  env,
	}
	itemGroup := group.Group("/api/v1/orders")
	itemGroup.GET("/ordersettings", lc.Fetch)
	itemGroup.GET("/ordersetting", lc.FetchByID)
	itemGroup.POST("/ordersetting", lc.Create)
	itemGroup.POST("/ordersettings", lc.CreateMany)
	itemGroup.PUT("/ordersetting", lc.Update)
	itemGroup.DELETE("/ordersetting", lc.Delete)
}
