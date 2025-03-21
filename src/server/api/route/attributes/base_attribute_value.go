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

	group.GET("/base_attribute_values", lc.Fetch)
	group.GET("/base_attribute_value", lc.FetchByID)
	group.POST("/base_attribute_value", lc.Create)
	group.PUT("/base_attribute_value", lc.Update)
	group.DELETE("/base_attribute_value", lc.Delete)
}
