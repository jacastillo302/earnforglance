package route

import (
	"time"

	controller "earnforglance/server/api/controller/discounts"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/discounts"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/discounts"
	usecase "earnforglance/server/usecase/discounts"

	"github.com/gin-gonic/gin"
)

func DiscountRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewDiscountRepository(db, domain.CollectionDiscount)
	lc := &controller.DiscountController{
		DiscountUsecase: usecase.NewDiscountUsecase(ur, timeout),
		Env:             env,
	}

	group.GET("/discounts", lc.Fetch)
	group.GET("/discount", lc.FetchByID)
	group.POST("/discount", lc.Create)
	group.PUT("/discount", lc.Update)
	group.DELETE("/discount", lc.Delete)
}
