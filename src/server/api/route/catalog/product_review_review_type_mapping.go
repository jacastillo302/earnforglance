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

func ProductReviewReviewTypeMappingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductReviewReviewTypeMappingRepository(db, domain.CollectionProductReviewReviewTypeMapping)
	lc := &controller.ProductReviewReviewTypeMappingController{
		ProductReviewReviewTypeMappingUsecase: usecase.NewProductReviewReviewTypeMappingUsecase(ur, timeout),
		Env:                                   env,
	}

	group.GET("/product_review_review_type_mappings", lc.Fetch)
	group.GET("/product_review_review_type_mapping", lc.FetchByID)
	group.POST("/product_review_review_type_mapping", lc.Create)
	group.POST("/product_review_review_type_mappings", lc.CreateMany)
	group.PUT("/product_review_review_type_mapping", lc.Update)
	group.DELETE("/product_review_review_type_mapping", lc.Delete)
}
