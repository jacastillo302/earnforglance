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

	group.GET("/apisettings", lc.Fetch)
	group.GET("/apisetting", lc.FetchByID)
	group.POST("/apisetting", lc.Create)
	Group.POST("/apisettings", lc.CreateMany)
	group.PUT("/apisetting", lc.Update)
	group.DELETE("apisetting", lc.Delete)
}
