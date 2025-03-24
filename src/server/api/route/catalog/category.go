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

	group.GET("/categories", lc.Fetch)
	group.GET("/category", lc.FetchByID)
	group.POST("/category", lc.Create)
	group.POST("/categories", lc.CreateMany)
	group.PUT("/category", lc.Update)
	group.DELETE("/category", lc.Delete)
}
