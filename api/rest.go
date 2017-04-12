package api

import gin "gopkg.in/gin-gonic/gin.v1"

//InstallRestAPI "instala" e sobe o servico de rest
func InstallRestAPI() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	InstallV1(router)
	router.Run(":3000")
}
