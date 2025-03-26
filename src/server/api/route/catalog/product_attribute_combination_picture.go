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

func ProductAttributeCombinationPictureRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductAttributeCombinationPictureRepository(db, domain.CollectionProductAttributeCombinationPicture)
	lc := &controller.ProductAttributeCombinationPictureController{
		ProductAttributeCombinationPictureUsecase: usecase.NewProductAttributeCombinationPictureUsecase(ur, timeout),
		Env: env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/product_attribute_combination_pictures", lc.Fetch)
	itemGroup.GET("/product_attribute_combination_picture", lc.FetchByID)
	itemGroup.POST("/product_attribute_combination_picture", lc.Create)
	itemGroup.POST("/product_attribute_combination_pictures", lc.CreateMany)
	itemGroup.PUT("/product_attribute_combination_picture", lc.Update)
	itemGroup.DELETE("/product_attribute_combination_picture", lc.Delete)
}
