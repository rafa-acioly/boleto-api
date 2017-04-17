package api

import (
	"bitbucket.org/mundipagg/boletoapi/config"
	gin "gopkg.in/gin-gonic/gin.v1"
)

//InstallRestAPI "instala" e sobe o servico de rest
func InstallRestAPI() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(ReturnHeaders())
	if config.Get().EnableRequestLog {
		router.Use(RequestLogger())
	}
	if config.Get().EnablePrintRequest {
		router.Use(gin.Logger())
	}
	InstallV1(router)
	router.Run(config.Get().APIPort)
}

type errorResponse struct {
	Code    string
	Message string
}
