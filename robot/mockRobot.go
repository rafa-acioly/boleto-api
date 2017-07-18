package robot

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http/httputil"

	"strings"

	gin "gopkg.in/gin-gonic/gin.v1"
)

//mockRobot sobe uma aplicação web para mockar a integração com os Bancos
func mockRobot() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.POST("/oauth/token", authBB)
	router.POST("/registrarBoleto", registerBoletoBB)
	router.POST("/caixa/registrarBoleto", registerBoletoCaixa)
	router.POST("/citi/registrarBoleto", registerBoletoCiti)

	router.Run(":4000")
}
func dumpRequest(c *gin.Context) {
	b, _ := httputil.DumpRequest(c.Request, true)
	fmt.Println("----------------------------------------------------------")
	fmt.Println(string(b))
	fmt.Println("----------------------------------------------------------")
}
func authBB(c *gin.Context) {

	const tok = `{
					"access_token":"Yemfehvhs9PmpKLcJgSLzjfjSxHj4QREdKcrhHbM_oivhlYXmOIPn5j2Tp6FdZFmbIzCxVN-SdxnyoGT7cE5xg.AcQSHv1xiN0uaZ-efTafZIWPNenudJn9eU54TUAfiR0ff8RRE9thqgdRb2gm9t_uTREmOsBOz9jvQySnsPBqbfptOqz9-O_63c-LQq2ogxbu7iet-6te8V28gfjOVePnr87yIK8ueATW9ulb7jytYRYJd7CuZXF3PyD763tI2ykX-PNm2LAClqpAU-WAORQ_2OSLo5ElwPS_MgVAqvXm_n1PX0wPazW-YlwSvoYr9pYabiBAOCW4KkiZva0hRver7AMWlkP2t2M_wttG6wv1V8szty2Lb9oyGDL-cdfdit4rHFgpXp9dzG3qFS5qWxtO5tnQc3sVBBybDNINOmMlOaxKrsrNtQl5ncELy6jozyrPS-Yb3JhlvVaj3IDHq599bd30G8JMjDsGSc9wuws7Ws9tuUbOTiS-d2TaGOqTjKs.Mvco1yM5ErKRswd11TBjFYr-zX0kor4y1EDyq52M_ew874ifju5PaU-G5btc2zMhUcltTIVY1sJlTf2rqXKaPQ",
					"token_type":"Bearer",
					"refresh_token":"eyJhbGciOiJBMTI4S1ciLCJlbmMiOiJBMTI4Q0JDLUhTMjU2In0.4PdGzdqXluMJ67StpSmi5Ds5rWUXiLKvhYvZh_HR8DAjBt361RaGdw.Gcg48k3Omleobjs-c5J_mw.CtxZAiHOf_oA3c4uPKzgGesG6V-Y9QzFhJh8ww262jI-GQL2S6YqWe1ucrJ9oY_hrST05Y1ns7rTZJkGluDBscNtE3mIuv-WkCykHUDlor2gevZjxUApj98mJIKeFqfaeIGnXZpyeQBpPXAcCIELIjUN4CAWm99ed72DCcCWiPbO3v2smSQVLX04ESKqTbnjRyHQLiHGm8jP4PnOFIafdBrnRSfhsqIggJCZYNfIC1aRIrDnTSDiTBdx1vEruLOCFIOv9z4pqySPbImzC3Uxv9UxNDKvEa11TGoVYlnAx62_8d7pFAC8IeDwXNuaRzFklyDWZCMNtFl0pEB1bqh3mN6QdeQE2sfsoMhyif9iXqcFnUJvFAu4Oj981M_Vyh2GW7VTAvs67sw27xvCS1diJZGNLR_O09WEjn529MZGyT_4oWqmlVTb-a6dflFWwdI3DhsusgvT6pK_ja-eIXq5pw.o50PzlpZnNnks17cNsaKog","expires_in":1200
				}`
	c.Data(200, "text/json", []byte(tok))
}

func registerBoletoBB(c *gin.Context) {

	sData := `
		<?xml version="1.0" encoding="UTF-8"?>
		<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
		<SOAP-ENV:Body>
			<ns0:resposta xmlns:ns0="http://www.tibco.com/schemas/bws_registro_cbr/Recursos/XSD/Schema.xsd">
				<ns0:siglaSistemaMensagem />
				<ns0:codigoRetornoPrograma>0</ns0:codigoRetornoPrograma>
				<ns0:nomeProgramaErro />
				<ns0:textoMensagemErro />
				<ns0:numeroPosicaoErroPrograma>0</ns0:numeroPosicaoErroPrograma>
				<ns0:codigoTipoRetornoPrograma>0</ns0:codigoTipoRetornoPrograma>
				<ns0:textoNumeroTituloCobrancaBb>00010140510000066673</ns0:textoNumeroTituloCobrancaBb>
				<ns0:numeroCarteiraCobranca>17</ns0:numeroCarteiraCobranca>
				<ns0:numeroVariacaoCarteiraCobranca>19</ns0:numeroVariacaoCarteiraCobranca>
				<ns0:codigoPrefixoDependenciaBeneficiario>3851</ns0:codigoPrefixoDependenciaBeneficiario>
				<ns0:numeroContaCorrenteBeneficiario>8570</ns0:numeroContaCorrenteBeneficiario>
				<ns0:codigoCliente>932131545</ns0:codigoCliente>
				<ns0:linhaDigitavel>00190000090101405100500066673179971340000010000</ns0:linhaDigitavel>
				<ns0:codigoBarraNumerico>00199713400000100000000001014051000006667317</ns0:codigoBarraNumerico>
				<ns0:codigoTipoEnderecoBeneficiario>0</ns0:codigoTipoEnderecoBeneficiario>
				<ns0:nomeLogradouroBeneficiario>Cliente nao informado.</ns0:nomeLogradouroBeneficiario>
				<ns0:nomeBairroBeneficiario />
				<ns0:nomeMunicipioBeneficiario />
				<ns0:codigoMunicipioBeneficiario>0</ns0:codigoMunicipioBeneficiario>
				<ns0:siglaUfBeneficiario />
				<ns0:codigoCepBeneficiario>0</ns0:codigoCepBeneficiario>
				<ns0:indicadorComprovacaoBeneficiario />
				<ns0:numeroContratoCobranca>17414296</ns0:numeroContratoCobranca>
			</ns0:resposta>
		</SOAP-ENV:Body>
		</SOAP-ENV:Envelope>
	`

	sDataErr := `
		<?xml version="1.0" encoding="UTF-8"?>
		<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
		<SOAP-ENV:Body>
			<ns0:resposta xmlns:ns0="http://www.tibco.com/schemas/bws_registro_cbr/Recursos/XSD/Schema.xsd">
				<ns0:siglaSistemaMensagem />
				<ns0:codigoRetornoPrograma>ER500</ns0:codigoRetornoPrograma>
				<ns0:nomeProgramaErro>Nome Programa ERRO</ns0:nomeProgramaErro>
				<ns0:textoMensagemErro>Falha ao registrar Boleto</ns0:textoMensagemErro>
				<ns0:numeroPosicaoErroPrograma>0</ns0:numeroPosicaoErroPrograma>
				<ns0:codigoTipoRetornoPrograma>0</ns0:codigoTipoRetornoPrograma>
				<ns0:textoNumeroTituloCobrancaBb>00010140510000066673</ns0:textoNumeroTituloCobrancaBb>
				<ns0:numeroCarteiraCobranca>17</ns0:numeroCarteiraCobranca>
				<ns0:numeroVariacaoCarteiraCobranca>19</ns0:numeroVariacaoCarteiraCobranca>
				<ns0:codigoPrefixoDependenciaBeneficiario>3851</ns0:codigoPrefixoDependenciaBeneficiario>
				<ns0:numeroContaCorrenteBeneficiario>8570</ns0:numeroContaCorrenteBeneficiario>
				<ns0:codigoCliente>932131545</ns0:codigoCliente>
				<ns0:linhaDigitavel>00190000090101405100500066673179971340000010000</ns0:linhaDigitavel>
				<ns0:codigoBarraNumerico>00199713400000100000000001014051000006667317</ns0:codigoBarraNumerico>
				<ns0:codigoTipoEnderecoBeneficiario>0</ns0:codigoTipoEnderecoBeneficiario>
				<ns0:nomeLogradouroBeneficiario>Cliente nao informado.</ns0:nomeLogradouroBeneficiario>
				<ns0:nomeBairroBeneficiario />
				<ns0:nomeMunicipioBeneficiario />
				<ns0:codigoMunicipioBeneficiario>0</ns0:codigoMunicipioBeneficiario>
				<ns0:siglaUfBeneficiario />
				<ns0:codigoCepBeneficiario>0</ns0:codigoCepBeneficiario>
				<ns0:indicadorComprovacaoBeneficiario />
				<ns0:numeroContratoCobranca>17414296</ns0:numeroContratoCobranca>
			</ns0:resposta>
		</SOAP-ENV:Body>
		</SOAP-ENV:Envelope>
	`

	d, _ := ioutil.ReadAll(c.Request.Body)
	xml := string(d)
	if strings.Contains(xml, "<sch:valorOriginalTitulo>2.00</sch:valorOriginalTitulo>") {
		c.Data(200, "text/xml", []byte(sData))
	} else {
		c.Data(200, "text/xml", []byte(sDataErr))
	}
}

func registerBoletoCaixa(c *gin.Context) {
	sData := `
<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
   <soapenv:Body>
      <manutencaocobrancabancaria:SERVICO_SAIDA xmlns:manutencaocobrancabancaria="http://caixa.gov.br/sibar/manutencao_cobranca_bancaria/boleto/externo" xmlns:sibar_base="http://caixa.gov.br/sibar">
         <sibar_base:HEADER>
            <VERSAO>1.0</VERSAO>
            <AUTENTICACAO>LvWr1op5Ayibn6jsCQ3/2bW4KwThVAlLK5ftxABlq20=</AUTENTICACAO>
            <USUARIO_SERVICO>SGCBS01D</USUARIO_SERVICO>
            <OPERACAO>INCLUI_BOLETO</OPERACAO>
            <SISTEMA_ORIGEM>SIGCB</SISTEMA_ORIGEM>
            <UNIDADE>1679</UNIDADE>
            <DATA_HORA>20170718150257</DATA_HORA>
         </sibar_base:HEADER>
         <COD_RETORNO>00</COD_RETORNO>
         <ORIGEM_RETORNO>MANUTENCAO_COBRANCA_BANCARIA</ORIGEM_RETORNO>
         <MSG_RETORNO />
         <DADOS>
            <CONTROLE_NEGOCIAL>
               <ORIGEM_RETORNO>SIGCB</ORIGEM_RETORNO>
               <COD_RETORNO>0</COD_RETORNO>
               <MENSAGENS>
                  <RETORNO>(0) OPERACAO EFETUADA</RETORNO>
               </MENSAGENS>
            </CONTROLE_NEGOCIAL>
            <INCLUI_BOLETO>
               <CODIGO_BARRAS>10493726700000010002006561000100040992226984</CODIGO_BARRAS>
               <LINHA_DIGITAVEL>10492006506100010004209922269841372670000001000</LINHA_DIGITAVEL>
               <NOSSO_NUMERO>14000000099222698</NOSSO_NUMERO>
               <URL>https://200.201.168.67:8010/ecobranca/SIGCB/imprimir/0200656/14000000099222698</URL>
            </INCLUI_BOLETO>
         </DADOS>
      </manutencaocobrancabancaria:SERVICO_SAIDA>
   </soapenv:Body>
</soapenv:Envelope>
	`

	sDataErr := `
<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
    <soapenv:Body>
        <manutencaocobrancabancaria:SERVICO_SAIDA xmlns:manutencaocobrancabancaria="http://caixa.gov.br/sibar/manutencao_cobranca_bancaria/boleto/externo" xmlns:sibar_base="http://caixa.gov.br/sibar">
            <sibar_base:HEADER>
                <VERSAO>1.0</VERSAO>
                <AUTENTICACAO>LvWr1op5Ayibn6jsCQ3/2bW4KwThVAlLK5ftxABlq201=</AUTENTICACAO>
                <USUARIO_SERVICO>SGCBS01D</USUARIO_SERVICO>
                <OPERACAO>INCLUI_BOLETO</OPERACAO>
                <SISTEMA_ORIGEM>SIGCB</SISTEMA_ORIGEM>
                <UNIDADE>1679</UNIDADE>
                <DATA_HORA>20170503102800</DATA_HORA>
            </sibar_base:HEADER>
            <COD_RETORNO>00</COD_RETORNO>
            <ORIGEM_RETORNO>MANUTENCAO_COBRANCA_BANCARIA</ORIGEM_RETORNO>
            <MSG_RETORNO></MSG_RETORNO>
            <DADOS>
                <CONTROLE_NEGOCIAL>
                    <ORIGEM_RETORNO>SIGCB</ORIGEM_RETORNO>
                    <COD_RETORNO>1</COD_RETORNO>
                    <MENSAGENS>
                        <RETORNO>(54) OPERACAO NAO PERMITIDA - HASH DIVERGENTE</RETORNO>
                    </MENSAGENS>
                </CONTROLE_NEGOCIAL>
            </DADOS>
        </manutencaocobrancabancaria:SERVICO_SAIDA>
    </soapenv:Body>
</soapenv:Envelope>
	`
	d, _ := ioutil.ReadAll(c.Request.Body)
	xml := string(d)
	if strings.Contains(xml, "<VALOR>5.04</VALOR>") {
		c.AbortWithError(504, errors.New("Teste de Erro"))
	} else if strings.Contains(xml, "<VALOR>2.00</VALOR>") {
		c.Data(200, "text/xml", []byte(sData))
	} else {
		c.Data(200, "text/xml", []byte(sDataErr))
	}

}

func registerBoletoCiti(c *gin.Context) {
	sData := `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
   <soapenv:Header/>
   <soapenv:Body>
      <RegisterBoletoResponse>
         <actionCode>0</actionCode>
         <reasonMessage>OK</reasonMessage>
      </RegisterBoletoResponse>
   </soapenv:Body>
</soapenv:Envelope>
	`

	sDataErr := `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
   <soapenv:Header/>
   <soapenv:Body>
      <RegisterBoletoResponse>
         <actionCode>99</actionCode>
         <reasonMessage>Erro ao Registrar boleto</reasonMessage>
      </RegisterBoletoResponse>
   </soapenv:Body>
</soapenv:Envelope>
	`
	d, _ := ioutil.ReadAll(c.Request.Body)
	xml := string(d)
	if strings.Contains(xml, "<TitlAmt>504</TitlAmt>") {
		c.AbortWithError(504, errors.New("Teste de Erro"))
	} else if strings.Contains(xml, "<TitlAmt>200</TitlAmt>") {
		c.Data(200, "text/xml", []byte(sData))
	} else {
		c.Data(200, "text/xml", []byte(sDataErr))
	}

}
