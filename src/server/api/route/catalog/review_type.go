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

func ReviewTypeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewReviewTypeRepository(db, domain.CollectionReviewType)
	lc := &controller.ReviewTypeController{
		ReviewTypeUsecase: usecase.NewReviewTypeUsecase(ur, timeout),
		Env:               env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/review_types", lc.Fetch)
	itemGroup.GET("/review_type", lc.FetchByID)
	itemGroup.POST("/review_type", lc.Create)
	itemGroup.POST("/review_types", lc.CreateMany)
	itemGroup.PUT("/review_type", lc.Update)
	itemGroup.DELETE("/review_type", lc.Delete)
}
