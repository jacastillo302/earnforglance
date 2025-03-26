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

func ProductReviewReviewTypeMappingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductReviewReviewTypeMappingRepository(db, domain.CollectionProductReviewReviewTypeMapping)
	lc := &controller.ProductReviewReviewTypeMappingController{
		ProductReviewReviewTypeMappingUsecase: usecase.NewProductReviewReviewTypeMappingUsecase(ur, timeout),
		Env:                                   env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/product_review_review_type_mappings", lc.Fetch)
	itemGroup.GET("/product_review_review_type_mapping", lc.FetchByID)
	itemGroup.POST("/product_review_review_type_mapping", lc.Create)
	itemGroup.POST("/product_review_review_type_mappings", lc.CreateMany)
	itemGroup.PUT("/product_review_review_type_mapping", lc.Update)
	itemGroup.DELETE("/product_review_review_type_mapping", lc.Delete)
}
