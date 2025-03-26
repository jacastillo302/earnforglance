package route

import (
	"time"

	controller "earnforglance/server/api/controller/affiliate"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/affiliate"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/affiliate"
	usecase "earnforglance/server/usecase/affiliate"

	"github.com/gin-gonic/gin"
)

func AffiliateRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewAffiliateRepository(db, domain.CollectionAffiliate)
	lc := &controller.AffiliateController{
		AffiliateUsecase: usecase.NewAffiliateUsecase(ur, timeout),
		Env:              env,
	}

	itemGroup := group.Group("/api/v1/affiliate")

	itemGroup.GET("/affiliates", lc.Fetch)
	itemGroup.GET("/affiliate", lc.FetchByID)
	itemGroup.POST("/affiliate", lc.Create)
	itemGroup.POST("/affiliates", lc.CreateMany)
	itemGroup.PUT("/affiliate", lc.Update)
	itemGroup.DELETE("/affiliate", lc.Delete)
}
