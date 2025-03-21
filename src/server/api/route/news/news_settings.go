package route

import (
	"time"

	controller "earnforglance/server/api/controller/news"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/news"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/news"
	usecase "earnforglance/server/usecase/news"

	"github.com/gin-gonic/gin"
)

func NewsSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewNewsSettingsRepository(db, domain.CollectionNewsSettings)
	lc := &controller.NewsSettingsController{
		NewsSettingsUsecase: usecase.NewNewsSettingsUsecase(ur, timeout),
		Env:                 env,
	}

	group.GET("/news_settingss", lc.Fetch)
	group.GET("/news_settings", lc.FetchByID)
	group.POST("/news_settings", lc.Create)
	group.PUT("/news_settings", lc.Update)
	group.DELETE("news_settings", lc.Delete)
}
