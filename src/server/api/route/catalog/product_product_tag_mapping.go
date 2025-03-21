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

func ProductProductTagMappingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductProductTagMappingRepository(db, domain.CollectionProductProductTagMapping)
	lc := &controller.ProductProductTagMappingController{
		ProductProductTagMappingUsecase: usecase.NewProductProductTagMappingUsecase(ur, timeout),
		Env:                             env,
	}

	group.GET("/product_product_tag_mappings", lc.Fetch)
	group.GET("/product_product_tag_mapping", lc.FetchByID)
	group.POST("/product_product_tag_mapping", lc.Create)
	group.PUT("/product_product_tag_mapping", lc.Update)
	group.DELETE("/product_product_tag_mapping", lc.Delete)
}
