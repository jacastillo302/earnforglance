package route

import (
	"time"

	controller "earnforglance/server/api/controller/common"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/common"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/common"
	usecase "earnforglance/server/usecase/common"

	"github.com/gin-gonic/gin"
)

func PdfSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPdfSettingsRepository(db, domain.CollectionPdfSettings)
	lc := &controller.PdfSettingsController{
		PdfSettingsUsecase: usecase.NewPdfSettingsUsecase(ur, timeout),
		Env:                env,
	}

	itemGroup := group.Group("/api/v1/common")
	itemGroup.GET("/pdf_settings", lc.Fetch)
	itemGroup.GET("/pdf_setting", lc.FetchByID)
	itemGroup.POST("/pdf_setting", lc.Create)
	itemGroup.POST("/pdf_settings", lc.CreateMany)
	itemGroup.PUT("/pdf_setting", lc.Update)
	itemGroup.DELETE("/pdf_setting", lc.Delete)
}
