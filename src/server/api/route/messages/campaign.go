package route

import (
	"time"

	controller "earnforglance/server/api/controller/messages"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/messages"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/messages"
	usecase "earnforglance/server/usecase/messages"

	"github.com/gin-gonic/gin"
)

func CampaignRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCampaignRepository(db, domain.CollectionCampaign)
	lc := &controller.CampaignController{
		CampaignUsecase: usecase.NewCampaignUsecase(ur, timeout),
		Env:             env,
	}
	itemGroup := group.Group("/api/v1/messages")
	itemGroup.GET("/campaigns", lc.Fetch)
	itemGroup.GET("/campaign", lc.FetchByID)
	itemGroup.POST("/campaign", lc.Create)
	itemGroup.POST("/campaigns", lc.CreateMany)
	itemGroup.PUT("/campaign", lc.Update)
	itemGroup.DELETE("/campaign", lc.Delete)
}
