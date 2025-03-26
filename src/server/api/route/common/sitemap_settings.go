package route

import (
	"time"

	controller "earnforglance/server/api/controller/common"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/common"

	repository "earnforglance/server/repository/common"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/common"

	"github.com/gin-gonic/gin"
)

func SitemapSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSitemapSettingsRepository(db, domain.CollectionSitemapSettings)
	lc := &controller.SitemapSettingsController{
		SitemapSettingsUsecase: usecase.NewSitemapSettingsUsecase(ur, timeout),
		Env:                    env,
	}

	itemGroup := group.Group("/api/v1/common")
	itemGroup.GET("/sitemap_settings", lc.Fetch)
	itemGroup.GET("/sitemap_setting", lc.FetchByID)
	itemGroup.POST("/sitemap_setting", lc.Create)
	itemGroup.POST("/sitemap_settings", lc.CreateMany)
	itemGroup.PUT("/sitemap_setting", lc.Update)
	itemGroup.DELETE("/sitemap_setting", lc.Delete)
}
