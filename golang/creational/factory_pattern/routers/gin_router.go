package routers

import (
	"factory-pattern/config"
	"factory-pattern/controller"
	"factory-pattern/logger"
	middlewarePkg "factory-pattern/middleware"
	"factory-pattern/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ginMiddleware middlewarePkg.IMiddleware[gin.HandlerFunc]

type ginRouter struct {
	ginEngine *gin.Engine
	env       config.IConfig
	logObj    logger.ILogger
}

func (router *ginRouter) StartApp() {

	router.ginEngine.Run(":8080")
}

func InitGinRouters(jobController controller.IJobController, logObj logger.ILogger) IRouter {
	ginEngine := gin.Default()
	routerGroup := getInitialRouteGroup(ginEngine)

	routerGroup.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Jobs api is working",
		})
	})

	routerGroup.POST("/getJobs", func(ginContext *gin.Context) {
		var request models.JobFilter
		if err := ginContext.ShouldBind(&request); err != nil {
			logObj.Printf("wrong query paramater %+v", err)
		}
		response, err := jobController.GetJobs(&request)
		if err != nil {
			logObj.Printf("error in getting jobs %+v", err)
			ginContext.JSON(http.StatusBadRequest, err)
			ginContext.Next()
		}
		ginContext.JSON(http.StatusOK, response)

	})

	routerGroup.GET("/getJobDetail/:jobId", func(ginContext *gin.Context) {
		jobId, _ := ginContext.Params.Get("jobId")
		response, err := jobController.GetJobDetail(jobId)
		if err != nil {
			logObj.Printf("error in getting jobs %+v", err)
			ginContext.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
			return
		}
		ginContext.JSON(http.StatusOK, response)
	})

	routerGroup.GET("/courses/:keywords", func(ginContext *gin.Context) {
		keywords, _ := ginContext.Params.Get("keywords")
		response, err := jobController.GetCourses(keywords)
		if err != nil {
			logObj.Printf("error in getting courses %+v", err)
			ginContext.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
			return
		}
		ginContext.JSON(http.StatusOK, response)
	})

	return &ginRouter{
		ginEngine: ginEngine,
		env:       config.GetConfig(),
		logObj:    logObj,
	}
}

func getInitialRouteGroup(ginEngine *gin.Engine) *gin.RouterGroup {
	return ginEngine.Group("/jobs")
}
