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

func CategoryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCategoryRepository(db, domain.CollectionCategory)
	lc := &controller.CategoryController{
		CategoryUsecase: usecase.NewCategoryUsecase(ur, timeout),
		Env:             env,
	}

	itemGroup := group.Group("/api/v1/catalog")

	itemGroup.GET("/categories", lc.Fetch)
	itemGroup.GET("/category", lc.FetchByID)
	itemGroup.POST("/category", lc.Create)
	itemGroup.POST("/categories", lc.CreateMany)
	itemGroup.PUT("/category", lc.Update)
	itemGroup.DELETE("/category", lc.Delete)
}
