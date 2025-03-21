package route

import (
	"time"

	controller "earnforglance/server/api/controller/gdpr"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/gdpr"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/gdpr"
	usecase "earnforglance/server/usecase/gdpr"

	"github.com/gin-gonic/gin"
)

func CustomerPermanentlyDeletedRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCustomerPermanentlyDeletedRepository(db, domain.CollectionCustomerPermanentlyDeleted)
	lc := &controller.CustomerPermanentlyDeletedController{
		CustomerPermanentlyDeletedUsecase: usecase.NewCustomerPermanentlyDeletedUsecase(ur, timeout),
		Env:                               env,
	}

	group.GET("/customer_permanently_deleteds", lc.Fetch)
	group.GET("/customer_permanently_deleted", lc.FetchByID)
	group.POST("/customer_permanently_deleted", lc.Create)
	group.PUT("/customer_permanently_deleted", lc.Update)
	group.DELETE("/customer_permanently_deleted", lc.Delete)
}
