package route

import (
	"time"

	controller "earnforglance/server/api/controller/catalog"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/catalog"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/catalog"
	usecase "earnforglance/server/usecase/catalog"

	"github.com/gin-gonic/gin"
)

func ProductEditorSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductEditorSettingsRepository(db, domain.CollectionProductEditorSettings)
	lc := &controller.ProductEditorSettingsController{
		ProductEditorSettingsUsecase: usecase.NewProductEditorSettingsUsecase(ur, timeout),
		Env:                          env,
	}

	group.GET("/product_editor_settings", lc.Fetch)
	group.GET("/product_editor_setting", lc.FetchByID)
	group.POST("/product_editor_setting", lc.Create)
	group.POST("/product_editor_settings", lc.CreateMany)
	group.PUT("/product_editor_setting", lc.Update)
	group.DELETE("/product_editor_setting", lc.Delete)
}
