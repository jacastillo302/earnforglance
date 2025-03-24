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

	group.GET("/product_attribute_combination_pictures", lc.Fetch)
	group.GET("/product_attribute_combination_picture", lc.FetchByID)
	group.POST("/product_attribute_combination_picture", lc.Create)
	group.POST("/product_attribute_combination_pictures", lc.CreateMany)
	group.PUT("/product_attribute_combination_picture", lc.Update)
	group.DELETE("/product_attribute_combination_picture", lc.Delete)
}
