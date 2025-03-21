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

func ProductAttributeMappingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductAttributeMappingRepository(db, domain.CollectionProductAttributeMapping)
	lc := &controller.ProductAttributeMappingController{
		ProductAttributeMappingUsecase: usecase.NewProductAttributeMappingUsecase(ur, timeout),
		Env:                            env,
	}

	group.GET("/product_attribute_mappings", lc.Fetch)
	group.GET("/product_attribute_mapping", lc.FetchByID)
	group.POST("/product_attribute_mapping", lc.Create)
	group.PUT("/product_attribute_mapping", lc.Update)
	group.DELETE("/product_attribute_mapping", lc.Delete)
}
