package route

import (
	"time"

	controller "earnforglance/server/api/controller/api"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/api"

	repository "earnforglance/server/repository/api"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/api"

	"github.com/gin-gonic/gin"
)

func ApiSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewApiSettingsRepository(db, domain.CollectionApiSettings)
	lc := &controller.ApiSettingsController{
		ApiSettingsUsecase: usecase.NewApiSettingsUsecase(ur, timeout),
		Env:                env,
	}

	Group := group.Group("/api/v1/api")

	Group.GET("/apisettings", lc.Fetch)
	Group.GET("/apisetting", lc.FetchByID)
	Group.POST("/apisetting", lc.Create)
	Group.POST("/apisettings", lc.CreateMany)
	Group.PUT("/apisetting", lc.Update)
	Group.DELETE("apisetting", lc.Delete)
}
