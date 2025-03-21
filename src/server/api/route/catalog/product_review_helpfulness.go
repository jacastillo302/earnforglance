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

func ProductReviewHelpfulnessRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductReviewHelpfulnessRepository(db, domain.CollectionProductReviewHelpfulness)
	lc := &controller.ProductReviewHelpfulnessController{
		ProductReviewHelpfulnessUsecase: usecase.NewProductReviewHelpfulnessUsecase(ur, timeout),
		Env:                             env,
	}

	group.GET("/product_review_helpfulnesses", lc.Fetch)
	group.GET("/product_review_helpfulness", lc.FetchByID)
	group.POST("/product_review_helpfulness", lc.Create)
	group.PUT("/product_review_helpfulness", lc.Update)
	group.DELETE("/product_review_helpfulness", lc.Delete)
}
