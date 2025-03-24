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

	group.GET("/ordersettings", lc.Fetch)
	group.GET("/ordersetting", lc.FetchByID)
	group.POST("/ordersetting", lc.Create)
	group.POST("/ordersettings", lc.CreateMany)
	group.PUT("/ordersetting", lc.Update)
	group.DELETE("/ordersetting", lc.Delete)
}
