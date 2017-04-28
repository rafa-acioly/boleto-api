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
	router.Use(executionController())
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
			Errors: models.NewEmptyErrorCollection(),
		}
		if e, ok := err.(models.IErrorResponse); ok {
			errResp.Errors.Append(e.ErrorCode(), e.Error())
			c.JSON(http.StatusBadRequest, errResp)
		} else if e, ok := err.(models.IErrorHTTP); ok {
			errResp.Errors.Append("MP500", e.Error())
			c.JSON(e.StatusCode(), errResp)
		} else {
			e := models.ErrorStatusHTTP{Code: 500, Message: err.Error()}
			errResp.Errors.Append("MP500", e.Error())
			c.JSON(e.StatusCode(), errResp)
		}
		return true
	}
	return false
}
