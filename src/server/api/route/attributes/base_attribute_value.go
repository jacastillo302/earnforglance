package route

import (
	"time"

	controller "earnforglance/server/api/controller/attributes"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/attributes"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/attributes"
	usecase "earnforglance/server/usecase/attributes"

	"github.com/gin-gonic/gin"
)

func BaseAttributeValueRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewBaseAttributeValueRepository(db, domain.CollectionBaseAttributeValue)
	lc := &controller.BaseAttributeValueController{
		BaseAttributeValueUsecase: usecase.NewBaseAttributeValueUsecase(ur, timeout),
		Env:                       env,
	}

	itemGroup := group.Group("/api/v1/attributes")

	itemGroup.GET("/base_attribute_values", lc.Fetch)
	itemGroup.GET("/base_attribute_value", lc.FetchByID)
	itemGroup.POST("/base_attribute_value", lc.Create)
	itemGroup.POST("/base_attribute_values", lc.CreateMany)
	itemGroup.PUT("/base_attribute_value", lc.Update)
	itemGroup.DELETE("/base_attribute_value", lc.Delete)
}
