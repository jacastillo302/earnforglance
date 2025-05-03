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

func PermisionRecordAttributeValueRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPermisionRecordAttributeValueRepository(db, domain.CollectionPermisionRecordAttributeValue)
	lc := &controller.PermisionRecordAttributeValueController{
		PermisionRecordAttributeValueUsecase: usecase.NewPermisionRecordAttributeValueUsecase(ur, timeout),
		Env:                                  env,
	}

	itemGroup := group.Group("/api/v1/attributes")
	itemGroup.GET("/permission_record_attribute_values", lc.Fetch)
	itemGroup.GET("/permission_record_attribute_value", lc.FetchByID)
	itemGroup.POST("/permission_record_attribute_value", lc.Create)
	itemGroup.POST("/permission_record_attribute_values", lc.CreateMany)
	itemGroup.PUT("/permission_record_attribute_value", lc.Update)
	itemGroup.DELETE("/permission_record_attribute_value", lc.Delete)
}
