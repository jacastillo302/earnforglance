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

	group.GET("/sitemap_xml_settings", lc.Fetch)
	group.GET("/sitemap_xml_setting", lc.FetchByID)
	group.POST("/sitemap_xml_setting", lc.Create)
	group.POST("/sitemap_xml_settings", lc.CreateMany)
	group.PUT("/sitemap_xml_setting", lc.Update)
	group.DELETE("sitemap_xml_setting", lc.Delete)
}
