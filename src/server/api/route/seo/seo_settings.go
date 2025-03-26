package route

import (
	"time"

	controller "earnforglance/server/api/controller/seo"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/seo"

	repository "earnforglance/server/repository/seo"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/seo"

	"github.com/gin-gonic/gin"
)

func SeoSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSeoSettingsRepository(db, domain.CollectionSeoSettings)
	lc := &controller.SeoSettingsController{
		SeoSettingsUsecase: usecase.NewSeoSettingsUsecase(ur, timeout),
		Env:                env,
	}
	itemGroup := group.Group("/api/v1/seo")
	itemGroup.GET("/seo_settings", lc.Fetch)
	itemGroup.GET("/seo_setting", lc.FetchByID)
	itemGroup.POST("/seo_setting", lc.Create)
	itemGroup.POST("/seo_settings", lc.CreateMany)
	itemGroup.PUT("/seo_setting", lc.Update)
	itemGroup.DELETE("/seo_setting", lc.Delete)
}
