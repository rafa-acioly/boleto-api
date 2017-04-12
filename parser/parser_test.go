package parser

import (
	"testing"
)

const xmlDoc = `
<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope 
    xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
    <SOAP-ENV:Body>
        <ns0:resposta 
            xmlns:ns0="http://www.tibco.com/schemas/bws_registro_cbr/Recursos/XSD/Schema.xsd">
            <ns0:siglaSistemaMensagem></ns0:siglaSistemaMensagem>
            <ns0:codigoRetornoPrograma>92</ns0:codigoRetornoPrograma>
            <ns0:nomeProgramaErro>CBRSR004</ns0:nomeProgramaErro>
            <ns0:textoMensagemErro>O Numero do Titulo informado nao esta disponivel.</ns0:textoMensagemErro>
            <ns0:numeroPosicaoErroPrograma>11</ns0:numeroPosicaoErroPrograma>
            <ns0:codigoTipoRetornoPrograma>0</ns0:codigoTipoRetornoPrograma>
            <ns0:textoNumeroTituloCobrancaBb></ns0:textoNumeroTituloCobrancaBb>
            <ns0:numeroCarteiraCobranca>17</ns0:numeroCarteiraCobranca>
            <ns0:numeroVariacaoCarteiraCobranca>19</ns0:numeroVariacaoCarteiraCobranca>
            <ns0:codigoPrefixoDependenciaBeneficiario>3851</ns0:codigoPrefixoDependenciaBeneficiario>
            <ns0:numeroContaCorrenteBeneficiario>8570</ns0:numeroContaCorrenteBeneficiario>
            <ns0:codigoCliente>932131545</ns0:codigoCliente>
            <ns0:linhaDigitavel></ns0:linhaDigitavel>
            <ns0:codigoBarraNumerico></ns0:codigoBarraNumerico>
            <ns0:codigoTipoEnderecoBeneficiario>0</ns0:codigoTipoEnderecoBeneficiario>
            <ns0:nomeLogradouroBeneficiario></ns0:nomeLogradouroBeneficiario>
            <ns0:nomeBairroBeneficiario></ns0:nomeBairroBeneficiario>
            <ns0:nomeMunicipioBeneficiario></ns0:nomeMunicipioBeneficiario>
            <ns0:codigoMunicipioBeneficiario>0</ns0:codigoMunicipioBeneficiario>
            <ns0:siglaUfBeneficiario></ns0:siglaUfBeneficiario>
            <ns0:codigoCepBeneficiario>0</ns0:codigoCepBeneficiario>
            <ns0:indicadorComprovacaoBeneficiario></ns0:indicadorComprovacaoBeneficiario>
            <ns0:numeroContratoCobranca>17414296</ns0:numeroContratoCobranca>
        </ns0:resposta>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>
`

func TestXmlParser(t *testing.T) {
	doc, err := ParseXML(xmlDoc)
	if err != nil {
		t.Fail()
	}
	translator := NewTranslatorMap()
	translator.AddRule(Rule{XMLQuery: "///ns0:resposta/ns0:codigoRetornoPrograma", MapKey: "returnCode"})
	values := ExtractValues(doc, translator)
	for _, rule := range translator.GetRules() {
		_, ok := values[rule.MapKey]
		if !ok {
			t.Fail()
		}
	}
}
