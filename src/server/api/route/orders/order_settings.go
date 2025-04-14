package route

import (
	"time"

	controller "earnforglance/server/api/controller/orders"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/orders"

	repository "earnforglance/server/repository/orders"
	"earnforglance/server/service/data/mongo"
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
	itemGroup.GET("/order_settings", lc.Fetch)
	itemGroup.GET("/order_setting", lc.FetchByID)
	itemGroup.POST("/order_setting", lc.Create)
	itemGroup.POST("/order_settings", lc.CreateMany)
	itemGroup.PUT("/order_setting", lc.Update)
	itemGroup.DELETE("/order_setting", lc.Delete)
}
