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

func ProductPictureRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductPictureRepository(db, domain.CollectionProductPicture)
	lc := &controller.ProductPictureController{
		ProductPictureUsecase: usecase.NewProductPictureUsecase(ur, timeout),
		Env:                   env,
	}

	group.GET("/product_pictures", lc.Fetch)
	group.GET("/product_picture", lc.FetchByID)
	group.POST("/product_picture", lc.Create)
	group.PUT("/product_picture", lc.Update)
	group.DELETE("/product_picture", lc.Delete)
}
