package route

import (
	"time"

	controller "earnforglance/server/api/controller/gdpr"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/gdpr"

	repository "earnforglance/server/repository/gdpr"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/gdpr"

	"github.com/gin-gonic/gin"
)

func CustomerPermanentlyDeletedRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCustomerPermanentlyDeletedRepository(db, domain.CollectionCustomerPermanentlyDeleted)
	lc := &controller.CustomerPermanentlyDeletedController{
		CustomerPermanentlyDeletedUsecase: usecase.NewCustomerPermanentlyDeletedUsecase(ur, timeout),
		Env:                               env,
	}
	itemGroup := group.Group("/api/v1/gdpr")
	itemGroup.GET("/customer_permanently_deleteds", lc.Fetch)
	itemGroup.GET("/customer_permanently_deleted", lc.FetchByID)
	itemGroup.POST("/customer_permanently_deleted", lc.Create)
	itemGroup.POST("/customer_permanently_deleteds", lc.CreateMany)
	itemGroup.PUT("/customer_permanently_deleted", lc.Update)
	itemGroup.DELETE("/customer_permanently_deleted", lc.Delete)
}
