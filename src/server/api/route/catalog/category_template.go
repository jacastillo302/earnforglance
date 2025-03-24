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

	group.GET("/category_templates", lc.Fetch)
	group.GET("/category_template", lc.FetchByID)
	group.POST("/category_template", lc.Create)
	group.POST("/category_templates", lc.CreateMany)
	group.PUT("/category_template", lc.Update)
	group.DELETE("/category_template", lc.Delete)
}
