package request

import (
	. "github.com/space-tech-dev/sms-service-notification/app/http"
	. "github.com/space-tech-dev/sms-service-notification/config"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/gin-gonic/gin"
)

type ParaLogin struct {
	UserName string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

func ValidateRequestLogin(context *gin.Context) {
	var form ParaLogin

	if err := context.ShouldBind(&form); err != nil {

		apiResponse := &APIResponse{}
		apiResponse.SetCode(API_ERR_CODE_REQUEST_PARAM_ERROR, API_RETURN_CODE_ERROR, "", "").SetData(object.HashMap{
			"message": err.Error(),
		}).ThrowJSONResponse(context)

	}
}
