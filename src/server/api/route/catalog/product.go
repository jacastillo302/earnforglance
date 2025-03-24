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

func ProductRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductRepository(db, domain.CollectionProduct)
	lc := &controller.ProductController{
		ProductUsecase: usecase.NewProductUsecase(ur, timeout),
		Env:            env,
	}

	group.GET("/products", lc.Fetch)
	group.GET("/product", lc.FetchByID)
	group.POST("/product", lc.Create)
	group.POST("/products", lc.CreateMany)
	group.PUT("/product", lc.Update)
	group.DELETE("/product", lc.Delete)
}
