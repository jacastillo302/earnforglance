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

func ProductTemplateRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductTemplateRepository(db, domain.CollectionProductTemplate)
	lc := &controller.ProductTemplateController{
		ProductTemplateUsecase: usecase.NewProductTemplateUsecase(ur, timeout),
		Env:                    env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/product_templates", lc.Fetch)
	itemGroup.GET("/product_template", lc.FetchByID)
	itemGroup.POST("/product_template", lc.Create)
	itemGroup.POST("/product_templates", lc.CreateMany)
	itemGroup.PUT("/product_template", lc.Update)
	itemGroup.DELETE("/product_template", lc.Delete)
}
