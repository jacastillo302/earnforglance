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

func ProductRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductRepository(db, domain.CollectionProduct)
	lc := &controller.ProductController{
		ProductUsecase: usecase.NewProductUsecase(ur, timeout),
		Env:            env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/products", lc.Fetch)
	itemGroup.GET("/product", lc.FetchByID)
	itemGroup.POST("/product", lc.Create)
	itemGroup.POST("/products", lc.CreateMany)
	itemGroup.PUT("/product", lc.Update)
	itemGroup.DELETE("/product", lc.Delete)
}
