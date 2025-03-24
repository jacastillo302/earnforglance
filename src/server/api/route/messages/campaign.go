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

	group.GET("/campaigns", lc.Fetch)
	group.GET("/campaign", lc.FetchByID)
	group.POST("/campaign", lc.Create)
	group.POST("/campaigns", lc.CreateMany)
	group.PUT("/campaign", lc.Update)
	group.DELETE("/campaign", lc.Delete)
}
