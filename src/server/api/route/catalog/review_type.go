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

	group.GET("/review_types", lc.Fetch)
	group.GET("/review_type", lc.FetchByID)
	group.POST("/review_type", lc.Create)
	group.PUT("/review_type", lc.Update)
	group.DELETE("/review_type", lc.Delete)
}
