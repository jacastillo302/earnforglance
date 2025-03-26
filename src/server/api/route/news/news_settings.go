package route

import (
	"time"

	controller "earnforglance/server/api/controller/news"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/news"

	repository "earnforglance/server/repository/news"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/news"

	"github.com/gin-gonic/gin"
)

func NewsSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewNewsSettingsRepository(db, domain.CollectionNewsSettings)
	lc := &controller.NewsSettingsController{
		NewsSettingsUsecase: usecase.NewNewsSettingsUsecase(ur, timeout),
		Env:                 env,
	}
	itemGroup := group.Group("/api/v1/news")
	itemGroup.GET("/news_settings", lc.Fetch)
	itemGroup.GET("/news_setting", lc.FetchByID)
	itemGroup.POST("/news_setting", lc.Create)
	itemGroup.POST("/news_settings", lc.CreateMany)
	itemGroup.PUT("/news_setting", lc.Update)
	itemGroup.DELETE("/news_setting", lc.Delete)
}
