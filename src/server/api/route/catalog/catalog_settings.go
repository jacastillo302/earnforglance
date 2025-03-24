package route

import (
	controller "earnforglance/server/api/controller/catalog"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/catalog"
	"time"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/catalog"
	usecase "earnforglance/server/usecase/catalog"

	"github.com/gin-gonic/gin"
)

func CatalogSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCatalogSettingsRepository(db, domain.CollectionCatalogSettings)
	lc := &controller.CatalogSettingsController{
		CatalogSettingsUsecase: usecase.NewCatalogSettingsUsecase(ur, timeout),
		Env:                    env,
	}

	group.GET("/catalog_settings", lc.Fetch)
	group.GET("/catalog_setting", lc.FetchByID)
	group.POST("/catalog_setting", lc.Create)
	group.POST("/catalog_settings", lc.CreateMany)
	group.PUT("/catalog_setting", lc.Update)
	group.DELETE("catalog_setting", lc.Delete)
}
