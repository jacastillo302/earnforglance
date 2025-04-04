package route

import (
	"time"

	controller "earnforglance/server/api/controller/customers"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/customers"

	repository "earnforglance/server/repository/customers"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/customers"

	"github.com/gin-gonic/gin"
)

func CustomerSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCustomerSettingsRepository(db, domain.CollectionCustomerSettings)
	lc := &controller.CustomerSettingsController{
		CustomerSettingsUsecase: usecase.NewCustomerSettingsUsecase(ur, timeout),
		Env:                     env,
	}

	Group := group.Group("/api/v1/customers")

	Group.GET("/customer_settings", lc.Fetch)
	Group.GET("/customer_setting", lc.FetchByID)
	Group.POST("/customer_setting", lc.Create)
	Group.POST("/customer_settings", lc.CreateMany)
	Group.PUT("/customer_setting", lc.Update)
	Group.DELETE("customer_setting", lc.Delete)
}
