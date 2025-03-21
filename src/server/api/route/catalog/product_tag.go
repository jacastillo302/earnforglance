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

func ProductTagRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductTagRepository(db, domain.CollectionProductTag)
	lc := &controller.ProductTagController{
		ProductTagUsecase: usecase.NewProductTagUsecase(ur, timeout),
		Env:               env,
	}

	group.GET("/product_tags", lc.Fetch)
	group.GET("/product_tag", lc.FetchByID)
	group.POST("/product_tag", lc.Create)
	group.PUT("/product_tag", lc.Update)
	group.DELETE("/product_tag", lc.Delete)
}
