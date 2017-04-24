package api

import (
	"net/http"

	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/models"
	gin "gopkg.in/gin-gonic/gin.v1"
)

//InstallRestAPI "instala" e sobe o servico de rest
func InstallRestAPI() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	if config.Get().EnablePrintRequest {
		//router.Use(gin.Logger())
	}
	InstallV1(router)
	router.GET("/boleto", getBoleto)
	router.Run(config.Get().APIPort)
}

func checkError(c *gin.Context, err error) bool {
	if err != nil {
		errResp := models.BoletoResponse{
			Errors: models.NewSingleErrorCollection("MP400", err.Error()),
		}
		c.JSON(http.StatusOK, errResp)
		return true
	}
	return false
}
