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

	group.GET("/product_attribute_value_pictures", lc.Fetch)
	group.GET("/product_attribute_value_picture", lc.FetchByID)
	group.POST("/product_attribute_value_picture", lc.Create)
	group.PUT("/product_attribute_value_picture", lc.Update)
	group.DELETE("/product_attribute_value_picture", lc.Delete)
}
