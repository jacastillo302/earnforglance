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

func ProductAttributeValuePictureRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductAttributeValuePictureRepository(db, domain.CollectionProductAttributeValuePicture)
	lc := &controller.ProductAttributeValuePictureController{
		ProductAttributeValuePictureUsecase: usecase.NewProductAttributeValuePictureUsecase(ur, timeout),
		Env:                                 env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/product_attribute_value_pictures", lc.Fetch)
	itemGroup.GET("/product_attribute_value_picture", lc.FetchByID)
	itemGroup.POST("/product_attribute_value_picture", lc.Create)
	itemGroup.POST("/product_attribute_value_pictures", lc.CreateMany)
	itemGroup.PUT("/product_attribute_value_picture", lc.Update)
	itemGroup.DELETE("/product_attribute_value_picture", lc.Delete)
}
