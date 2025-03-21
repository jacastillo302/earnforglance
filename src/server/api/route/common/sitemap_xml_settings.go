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

func SitemapXmlSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSitemapXmlSettingsRepository(db, domain.CollectionSitemapXmlSettings)
	lc := &controller.SitemapXmlSettingsController{
		SitemapXmlSettingsUsecase: usecase.NewSitemapXmlSettingsUsecase(ur, timeout),
		Env:                       env,
	}

	group.GET("/sitemap_xml_settingss", lc.Fetch)
	group.GET("/sitemap_xml_settings", lc.FetchByID)
	group.POST("/sitemap_xml_settings", lc.Create)
	group.PUT("/sitemap_xml_settings", lc.Update)
	group.DELETE("sitemap_xml_settings", lc.Delete)
}
