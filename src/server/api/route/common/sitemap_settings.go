package route

import (
	"time"

	controller "earnforglance/server/api/controller/common"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/common"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/common"
	usecase "earnforglance/server/usecase/common"

	"github.com/gin-gonic/gin"
)

func SitemapSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSitemapSettingsRepository(db, domain.CollectionSitemapSettings)
	lc := &controller.SitemapSettingsController{
		SitemapSettingsUsecase: usecase.NewSitemapSettingsUsecase(ur, timeout),
		Env:                    env,
	}

	group.GET("/sitemap_settings", lc.Fetch)
	group.GET("/sitemap_setting", lc.FetchByID)
	group.POST("/sitemap_setting", lc.Create)
	group.PUT("/sitemap_setting", lc.Update)
	group.DELETE("sitemap_setting", lc.Delete)
}
