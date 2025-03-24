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

func ProductCategoryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductCategoryRepository(db, domain.CollectionProductCategory)
	lc := &controller.ProductCategoryController{
		ProductCategoryUsecase: usecase.NewProductCategoryUsecase(ur, timeout),
		Env:                    env,
	}

	group.GET("/product_categories", lc.Fetch)
	group.GET("/product_category", lc.FetchByID)
	group.POST("/product_category", lc.Create)
	group.POST("/product_categories", lc.CreateMany)
	group.DELETE("/product_category", lc.Delete)
}
