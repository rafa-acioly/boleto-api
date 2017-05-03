package api

import (
	"net/http"

	"bitbucket.org/mundipagg/boletoapi/log"

	"bitbucket.org/mundipagg/boletoapi/config"

	"bitbucket.org/mundipagg/boletoapi/models"
	gin "gopkg.in/gin-gonic/gin.v1"
)

//InstallRestAPI "instala" e sobe o servico de rest
func InstallRestAPI() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(executionController())
	if config.Get().DevMode {
		router.Use(gin.Logger())
	}
	InstallV1(router)
	router.GET("/boleto", getBoleto)
	router.Run(config.Get().APIPort)

}

func checkError(c *gin.Context, err error, log *log.Log) bool {
	if err != nil {
		errResp := models.BoletoResponse{
			Errors: models.NewErrors(),
		}
		if e, ok := err.(models.IErrorResponse); ok {
			errResp.Errors.Append(e.ErrorCode(), e.Error())
			c.JSON(http.StatusBadRequest, errResp)
		} else if e, ok := err.(models.IServerError); ok {
			errResp.Errors.Append("MP500", "Erro interno")
			errResp.StatusCode = http.StatusInternalServerError
			log.Fatal(e.Error(), e.Message())
			c.JSON(http.StatusInternalServerError, errResp)
		} else {
			log.Fatal(err.Error(), "")
			errResp.Errors.Append("MP500", "Erro interno")
			c.JSON(http.StatusInternalServerError, errResp)
		}
		return true
	}
	return false
}
