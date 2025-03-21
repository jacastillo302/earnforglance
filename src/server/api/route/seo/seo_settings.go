package route

import (
	"time"

	controller "earnforglance/server/api/controller/seo"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/seo"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/seo"
	usecase "earnforglance/server/usecase/seo"

	"github.com/gin-gonic/gin"
)

func SeoSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSeoSettingsRepository(db, domain.CollectionSeoSettings)
	lc := &controller.SeoSettingsController{
		SeoSettingsUsecase: usecase.NewSeoSettingsUsecase(ur, timeout),
		Env:                env,
	}

	group.GET("/seo_settings", lc.Fetch)
	group.GET("/seo_setting", lc.FetchByID)
	group.POST("/seo_setting", lc.Create)
	group.PUT("/seo_setting", lc.Update)
	group.DELETE("/seo_setting", lc.Delete)
}
