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

	group.GET("/settings", lc.Fetch)
	group.GET("/setting", lc.FetchByID)
	group.POST("/setting", lc.Create)
	group.PUT("/setting", lc.Update)
	group.DELETE("setting", lc.Delete)
}
