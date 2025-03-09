package main

import (
	"factory-pattern/config"
	"factory-pattern/controller"
	"factory-pattern/logger"
	"factory-pattern/service"
	"fmt"
)

var envVariables config.IConfig
var jobDetailsService service.IJobService
var jobControllerObj controller.IJobController
var logObj logger.ILogger

func main() {
	initLogger()
	initConfig()
	fmt.Println(envVariables)
	// initRepository()
	// intitServices()
	// initControllers()

	// routers.InitGinRouters(jobControllerObj, logObj).StartApp()
}
func initLogger() {
	logObj = logger.InitConsoleLogger()
}
func initConfig() {
	envVariables = config.GetConfig()
}

func intitServices() {
	jobDetailsService = service.InitJobService()
}

func initControllers() {
	jobControllerObj = controller.InitJobController(jobDetailsService, logObj)
}
