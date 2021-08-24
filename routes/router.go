package routes

import (
	logger "github.com/space-tech-dev/sms-service-notification/loggerManager"
	UBT "github.com/ArtisanCloud/ubt-go"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitializeRoutes(router *gin.Engine) {

	Router = router
	Router.Use(UBT.GinEsLog(logger.UBTHandler))
	//println("%T", router)
	InitializeWebRoutes(router)
	InitializeAPIRoutes(router)
}
