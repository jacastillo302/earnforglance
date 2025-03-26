package route

import (
	"time"

	controller "earnforglance/server/api/controller/configuration"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/configuration"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/configuration"
	usecase "earnforglance/server/usecase/configuration"

	"github.com/gin-gonic/gin"
)

func SettingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSettingRepository(db, domain.CollectionSetting)
	lc := &controller.SettingController{
		SettingUsecase: usecase.NewSettingUsecase(ur, timeout),
		Env:            env,
	}

	itemGroup := group.Group("/api/v1/configuration")
	itemGroup.GET("/settings", lc.Fetch)
	itemGroup.GET("/setting", lc.FetchByID)
	itemGroup.POST("/setting", lc.Create)
	itemGroup.POST("/settings", lc.CreateMany)
	itemGroup.PUT("/setting", lc.Update)
	itemGroup.DELETE("/setting", lc.Delete)
}
