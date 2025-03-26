package route

import (
	"time"

	controller "earnforglance/server/api/controller/catalog"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/catalog"

	repository "earnforglance/server/repository/catalog"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/catalog"

	"github.com/gin-gonic/gin"
)

func ProductEditorSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductEditorSettingsRepository(db, domain.CollectionProductEditorSettings)
	lc := &controller.ProductEditorSettingsController{
		ProductEditorSettingsUsecase: usecase.NewProductEditorSettingsUsecase(ur, timeout),
		Env:                          env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/product_editor_settings", lc.Fetch)
	itemGroup.GET("/product_editor_setting", lc.FetchByID)
	itemGroup.POST("/product_editor_setting", lc.Create)
	itemGroup.POST("/product_editor_settings", lc.CreateMany)
	itemGroup.PUT("/product_editor_setting", lc.Update)
	itemGroup.DELETE("/product_editor_setting", lc.Delete)
}
