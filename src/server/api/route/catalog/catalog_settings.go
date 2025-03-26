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

	itemGroup := group.Group("/api/v1/catalog")

	itemGroup.GET("/catalog_settings", lc.Fetch)
	itemGroup.GET("/catalog_setting", lc.FetchByID)
	itemGroup.POST("/catalog_setting", lc.Create)
	itemGroup.POST("/catalog_settings", lc.CreateMany)
	itemGroup.PUT("/catalog_setting", lc.Update)
	itemGroup.DELETE("/catalog_setting", lc.Delete)
}
