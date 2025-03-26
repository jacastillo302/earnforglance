package route

import (
	"time"

	controller "earnforglance/server/api/controller/seo"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/seo"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/seo"
	usecase "earnforglance/server/usecase/seo"

	"github.com/gin-gonic/gin"
)

func UrlRecordRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUrlRecordRepository(db, domain.CollectionUrlRecord)
	lc := &controller.UrlRecordController{
		UrlRecordUsecase: usecase.NewUrlRecordUsecase(ur, timeout),
		Env:              env,
	}
	itemGroup := group.Group("/api/v1/seo")
	itemGroup.GET("/url_records", lc.Fetch)
	itemGroup.GET("/url_record", lc.FetchByID)
	itemGroup.POST("/url_record", lc.Create)
	itemGroup.POST("/url_records", lc.CreateMany)
	itemGroup.PUT("/url_record", lc.Update)
	itemGroup.DELETE("/url_record", lc.Delete)
}
