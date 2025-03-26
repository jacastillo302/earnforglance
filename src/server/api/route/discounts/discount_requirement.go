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

func DiscountRequirementRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewDiscountRequirementRepository(db, domain.CollectionDiscountRequirement)
	lc := &controller.DiscountRequirementController{
		DiscountRequirementUsecase: usecase.NewDiscountRequirementUsecase(ur, timeout),
		Env:                        env,
	}

	itemGroup := group.Group("/api/v1/discounts")
	itemGroup.GET("/discount_requirements", lc.Fetch)
	itemGroup.GET("/discount_requirement", lc.FetchByID)
	itemGroup.POST("/discount_requirement", lc.Create)
	itemGroup.POST("/discount_requirements", lc.CreateMany)
	itemGroup.PUT("/discount_requirement", lc.Update)
	itemGroup.DELETE("/discount_requirement", lc.Delete)
}
