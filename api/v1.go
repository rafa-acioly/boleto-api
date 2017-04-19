package api

import (
	"bitbucket.org/mundipagg/boletoapi/config"
	gin "gopkg.in/gin-gonic/gin.v1"
)

//InstallV1 instala a api versao 1
func InstallV1(router *gin.Engine) {
	v1 := router.Group("v1")
	v1.POST("/boleto/register", registerBoleto)
	v1.GET("/boleto", getBoleto)
	v1.GET("/info", func(c *gin.Context) {
		c.JSON(200, config.Get())
	})
}
