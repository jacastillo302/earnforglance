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

func ProductReviewRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductReviewRepository(db, domain.CollectionProductReview)
	lc := &controller.ProductReviewController{
		ProductReviewUsecase: usecase.NewProductReviewUsecase(ur, timeout),
		Env:                  env,
	}

	group.GET("/product_reviews", lc.Fetch)
	group.GET("/product_review", lc.FetchByID)
	group.POST("/product_review", lc.Create)
	group.PUT("/product_review", lc.Update)
	group.DELETE("/product_review", lc.Delete)
}
