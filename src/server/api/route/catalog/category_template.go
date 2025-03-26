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

func CategoryTemplateRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCategoryTemplateRepository(db, domain.CollectionCategoryTemplate)
	lc := &controller.CategoryTemplateController{
		CategoryTemplateUsecase: usecase.NewCategoryTemplateUsecase(ur, timeout),
		Env:                     env,
	}

	itemGroup := group.Group("/api/v1/catalog")

	itemGroup.GET("/category_templates", lc.Fetch)
	itemGroup.GET("/category_template", lc.FetchByID)
	itemGroup.POST("/category_template", lc.Create)
	itemGroup.POST("/category_templates", lc.CreateMany)
	itemGroup.PUT("/category_template", lc.Update)
	itemGroup.DELETE("/category_template", lc.Delete)
}
