package routes

import (
	"github.com/gin-gonic/gin"
	. "github.com/space-tech-dev/sms-service-notification/app/http/controllers"
	. "github.com/space-tech-dev/sms-service-notification/app/http/middleware"
	. "github.com/space-tech-dev/sms-service-notification/app/http/request"
)

func InitializeAPIRoutes(router *gin.Engine) {

	apiRouter := router.Group("/api")
	{
		apiRouter.Use(Maintenance, AuthAPI, AuthWeb)
		{
			// Handle the index route
			apiRouter.GET("/", APIGetHome)
			apiRouter.POST("/reservation/success", ValidatePushReservationSuccess, APIPushReservationSuccess)
			apiRouter.GET("/thirdPartyShop/list", APIGetThirdPartyShops)
			//apiRouter.POST("/reservation/success", ValidatePushReservationSuccess, APIPushReservationSuccess)
			// apiRouter.POST("/make", ValidateRequestMakeWelcome, ctlWelcome.APIMakeWelcome)
			// apiRouter.PUT("/somePut", putting)
			// apiRouter.DELETE("/someDelete", deleting)
			// apiRouter.PATCH("/somePatch", patching)
			// apiRouter.HEAD("/someHead", head)
			// apiRouter.OPTIONS("/someOptions", options)

		}
	}
}
