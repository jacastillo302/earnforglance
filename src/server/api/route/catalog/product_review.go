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

func ProductReviewRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductReviewRepository(db, domain.CollectionProductReview)
	lc := &controller.ProductReviewController{
		ProductReviewUsecase: usecase.NewProductReviewUsecase(ur, timeout),
		Env:                  env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/product_reviews", lc.Fetch)
	itemGroup.GET("/product_review", lc.FetchByID)
	itemGroup.POST("/product_review", lc.Create)
	itemGroup.POST("/product_reviews", lc.CreateMany)
	itemGroup.PUT("/product_review", lc.Update)
	itemGroup.DELETE("/product_review", lc.Delete)
}
