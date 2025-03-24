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

func TierPriceRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewTierPriceRepository(db, domain.CollectionTierPrice)
	lc := &controller.TierPriceController{
		TierPriceUsecase: usecase.NewTierPriceUsecase(ur, timeout),
		Env:              env,
	}

	group.GET("/tier_prices", lc.Fetch)
	group.GET("/tier_price", lc.FetchByID)
	group.POST("/tier_price", lc.Create)
	group.POST("/tier_prices", lc.CreateMany)
	group.PUT("/tier_price", lc.Update)
	group.DELETE("/tier_price", lc.Delete)
}
