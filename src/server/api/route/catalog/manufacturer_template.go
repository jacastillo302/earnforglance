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

	group.GET("/manufacturer_templates", lc.Fetch)
	group.GET("/manufacturer_template", lc.FetchByID)
	group.POST("/manufacturer_template", lc.Create)
	group.PUT("/manufacturer_template", lc.Update)
	group.DELETE("/manufacturer_template", lc.Delete)
}
