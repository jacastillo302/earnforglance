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

func NewAffiliateRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewAffiliateRepository(db, domain.CollectionAffiliate)
	lc := &controller.AffiliateController{
		AffiliateUsecase: usecase.NewAffiliateUsecase(ur, timeout),
		Env:              env,
	}
	group.GET("/affiliate", lc.FetchByID)
	group.GET("/affiliates", lc.Fetch)
	group.POST("/affiliate", lc.Create)
	group.PUT("/affiliate", lc.Update)
}
