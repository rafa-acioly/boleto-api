package api

import gin "gopkg.in/gin-gonic/gin.v1"

//Regista um boleto em um determinado banco
func registerBoleto(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "OK"})
}
