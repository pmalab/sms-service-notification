package main

import (
	service "github.com/space-tech-dev/sms-service-notification/app/services"
	"github.com/space-tech-dev/sms-service-notification/cache"
	"github.com/space-tech-dev/sms-service-notification/config"
	"github.com/space-tech-dev/sms-service-notification/database"
	logger "github.com/space-tech-dev/sms-service-notification/loggerManager"
	"github.com/space-tech-dev/sms-service-notification/resources/lang"
	tester "github.com/space-tech-dev/sms-service-notification/tests"
)
import _ "github.com/space-tech-dev/sms-service-notification/config"
import "github.com/space-tech-dev/sms-service-notification/routes"

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {

	// Initialize the global config
	envConfigName := "environment"
	dbConfigName := "database"
	cacheConfigName := "cache"
	logConfigName := "log"
	config.LoadEnvConfig(nil, &envConfigName, nil)
	config.LoadDatabaseConfig(nil, &dbConfigName, nil)
	config.LoadCacheConfig(nil, &cacheConfigName, nil)
	config.LoadVersion()
	config.LoadLogConfig(nil, &logConfigName, nil)

	// load locale
	lang.LoadLanguages()

	// setup ssh key path
	service.SetupSSHKeyPath(config.AppConfigure.SSH)

	// Initialize the cache
	_ = cache.SetupCache()

	// Initialize the database
	_ = database.SetupDatabase()

	// Initialize the Aliyun service
	_, _ = service.CreateClient(config.AppConfigure.AliYun)

	// Initialize the logger
	_ = logger.SetupLog()

	// Initialize the Logger
	tester.TestFun()




	// Router the router as the default one provided by Gin
	Router = gin.Default()

	// Initialize the routes
	routes.InitializeRoutes(Router)

	//templateParamJson := "{\"cardnumber\":\"201314570325\",\"deductamount\":\"520\"}"
	aliYunService := service.NewAliYunService()
	aliYunService.SendMsg("15732029254", "SPAC", "SMS_217418916", "")
	//aliYunService.SendMsg("15732029254", "SPACE", "SMS_211488651", templateParamJson)

	// Start serving the application
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080®")
	err := Router.Run(config.AppConfigure.Server.Host + ":" + config.AppConfigure.Server.Port)
	if err != nil {
		logger.Error("router error:", map[string]interface{}{
			"err": err,
		})
	}

}
