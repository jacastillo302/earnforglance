package route

import (
	"time"

	controller "earnforglance/server/api/controller/customers"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/customers"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/customers"
	usecase "earnforglance/server/usecase/customers"

	"github.com/gin-gonic/gin"
)

func BestCustomerReportLineRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewBestCustomerReportLineRepository(db, domain.CollectionBestCustomerReportLine)
	lc := &controller.BestCustomerReportLineController{
		BestCustomerReportLineUsecase: usecase.NewBestCustomerReportLineUsecase(ur, timeout),
		Env:                           env,
	}

	group.GET("/best_customer_report_lines", lc.Fetch)
	group.GET("/best_customer_report_line", lc.FetchByID)
	group.POST("/best_customer_report_line", lc.Create)
	group.PUT("/best_customer_report_line", lc.Update)
	group.DELETE("/best_customer_report_line", lc.Delete)
}
