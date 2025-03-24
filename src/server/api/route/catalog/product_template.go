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

func ProductTemplateRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductTemplateRepository(db, domain.CollectionProductTemplate)
	lc := &controller.ProductTemplateController{
		ProductTemplateUsecase: usecase.NewProductTemplateUsecase(ur, timeout),
		Env:                    env,
	}

	group.GET("/product_templates", lc.Fetch)
	group.GET("/product_template", lc.FetchByID)
	group.POST("/product_template", lc.Create)
	group.POST("/product_templates", lc.CreateMany)
	group.PUT("/product_template", lc.Update)
	group.DELETE("/product_template", lc.Delete)
}
