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

func ProductAttributeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductAttributeRepository(db, domain.CollectionProductAttribute)
	lc := &controller.ProductAttributeController{
		ProductAttributeUsecase: usecase.NewProductAttributeUsecase(ur, timeout),
		Env:                     env,
	}

	group.GET("/product_attributes", lc.Fetch)
	group.GET("/product_attribute", lc.FetchByID)
	group.POST("/product_attribute", lc.Create)
	group.POST("/product_attributes", lc.CreateMany)
	group.PUT("/product_attribute", lc.Update)
	group.DELETE("/product_attribute", lc.Delete)
}
