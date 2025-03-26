package route

import (
	"time"

	controller "earnforglance/server/api/controller/attributes"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/attributes"

	repository "earnforglance/server/repository/attributes"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/attributes"

	"github.com/gin-gonic/gin"
)

func BaseAttributeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewBaseAttributeRepository(db, domain.CollectionBaseAttribute)
	lc := &controller.BaseAttributeController{
		BaseAttributeUsecase: usecase.NewBaseAttributeUsecase(ur, timeout),
		Env:                  env,
	}

	group.Group("/attributes")

	group.GET("/base_attributes", lc.Fetch)
	group.GET("/base_attribute", lc.FetchByID)
	group.POST("/base_attribute", lc.Create)
	group.POST("/base_attributes", lc.CreateMany)
	group.PUT("/base_attribute", lc.Update)
	group.DELETE("/base_attribute", lc.Delete)
}
