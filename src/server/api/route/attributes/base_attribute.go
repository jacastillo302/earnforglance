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
	itemGroup := group.Group("/api/v1/attributes")
	itemGroup.GET("/base_attributes", lc.Fetch)
	itemGroup.GET("/base_attribute", lc.FetchByID)
	itemGroup.POST("/base_attribute", lc.Create)
	itemGroup.POST("/base_attributes", lc.CreateMany)
	itemGroup.PUT("/base_attribute", lc.Update)
	itemGroup.DELETE("/base_attribute", lc.Delete)
}
