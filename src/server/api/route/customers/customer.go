package route

import (
	"time"

	controller "earnforglance/server/api/controller/customers"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/customers"

	repository "earnforglance/server/repository/customers"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/customers"

	"github.com/gin-gonic/gin"
)

func CustomerRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCustomerRepository(db, domain.CollectionCustomer)
	lc := &controller.CustomerController{
		CustomerUsecase: usecase.NewCustomerUsecase(ur, timeout),
		Env:             env,
	}

	itemGroup := group.Group("/api/v1/customers")
	itemGroup.GET("/customers", lc.Fetch)
	itemGroup.GET("/customer", lc.FetchByID)
	itemGroup.POST("/customer", lc.Create)
	itemGroup.POST("/customers", lc.CreateMany)
	itemGroup.PUT("/customer", lc.Update)
	itemGroup.DELETE("/customer", lc.Delete)
}
