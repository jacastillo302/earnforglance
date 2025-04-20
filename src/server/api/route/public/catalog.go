package route

import (
	"time"

	controller "earnforglance/server/api/controller/public"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/public"

	repository "earnforglance/server/repository/public"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/public"

	"github.com/gin-gonic/gin"
)

func CatalogRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCatalogRepository(db, domain.CollectionUser)
	lc := &controller.CatalogController{
		CatalogUsecase: usecase.NewCatalogtUsecase(ur, timeout),
		Env:            env,
	}
	group.GET("/products", lc.GetProducts)
	group.GET("/categories", lc.GetCategories)
}
