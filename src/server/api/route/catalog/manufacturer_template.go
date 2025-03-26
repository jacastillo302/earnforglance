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

func ManufacturerTemplateRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewManufacturerTemplateRepository(db, domain.CollectionManufacturerTemplate)
	lc := &controller.ManufacturerTemplateController{
		ManufacturerTemplateUsecase: usecase.NewManufacturerTemplateUsecase(ur, timeout),
		Env:                         env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/manufacturer_templates", lc.Fetch)
	itemGroup.GET("/manufacturer_template", lc.FetchByID)
	itemGroup.POST("/manufacturer_template", lc.Create)
	itemGroup.POST("/manufacturer_templates", lc.CreateMany)
	itemGroup.PUT("/manufacturer_template", lc.Update)
	itemGroup.DELETE("/manufacturer_template", lc.Delete)
}
