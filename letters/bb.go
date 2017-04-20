package letters

import "bitbucket.org/mundipagg/boletoapi/parser"

/*
@author Philippe Moneda
@date 10/04/2017

Descreve o padrão de mensagem para Boletos do Banco do Brasil

Exemplo da Documentação:
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:sch="http://www.tibco.com/schemas/bws_registro_cbr/Recursos/XSD/Schema.xsd">
 <soapenv:Header/>
 <soapenv:Body>
<sch:requisicao>
 <sch:numeroConvenio>1014051</sch:numeroConvenio>
 <sch:numeroCarteira>17</sch:numeroCarteira>
 <sch:numeroVariacaoCarteira>19</sch:numeroVariacaoCarteira>
 <sch:codigoModalidadeTitulo>1</sch:codigoModalidadeTitulo>
 <sch:dataEmissaoTitulo>01.03.2017</sch:dataEmissaoTitulo>
 <sch:dataVencimentoTitulo>21.11.2017</sch:dataVencimentoTitulo>
 <sch:valorOriginalTitulo>30000</sch:valorOriginalTitulo>
 <sch:codigoTipoDesconto>1</sch:codigoTipoDesconto>
 <sch:dataDescontoTitulo>21.11.2016</sch:dataDescontoTitulo>
 <sch:percentualDescontoTitulo/>
 <sch:valorDescontoTitulo>10</sch:valorDescontoTitulo>
 <sch:valorAbatimentoTitulo/>
 <sch:quantidadeDiaProtesto>0</sch:quantidadeDiaProtesto>
 <sch:codigoTipoJuroMora>0</sch:codigoTipoJuroMora>
 <sch:percentualJuroMoraTitulo></sch:percentualJuroMoraTitulo>
 <sch:valorJuroMoraTitulo></sch:valorJuroMoraTitulo>
 <sch:codigoTipoMulta>2</sch:codigoTipoMulta>
 <sch:dataMultaTitulo>22.11.2017</sch:dataMultaTitulo>
 <sch:percentualMultaTitulo>10</sch:percentualMultaTitulo>
 <sch:valorMultaTitulo></sch:valorMultaTitulo>
 <sch:codigoAceiteTitulo>N</sch:codigoAceiteTitulo>
 <sch:codigoTipoTitulo>2</sch:codigoTipoTitulo>
 <sch:textoDescricaoTipoTitulo>DUPLICATA</sch:textoDescricaoTipoTitulo>
 <sch:indicadorPermissaoRecebimentoParcial>N</sch:indicadorPermissaoRecebimentoParcial>
 <sch:textoNumeroTituloBeneficiario>987654321987654</sch:textoNumeroTituloBeneficiario>
 <sch:textoCampoUtilizacaoBeneficiario/>
 <sch:codigoTipoContaCaucao>1</sch:codigoTipoContaCaucao>
 <sch:textoNumeroTituloCliente>00010140510000000630</sch:textoNumeroTituloCliente>
 <sch:textoMensagemBloquetoOcorrencia>Pagamento disponível até a data de vencimento
</sch:textoMensagemBloquetoOcorrencia>
 <sch:codigoTipoInscricaoPagador>2</sch:codigoTipoInscricaoPagador>
 <sch:numeroInscricaoPagador>73400584000166</sch:numeroInscricaoPagador>
 <sch:nomePagador>MERCADO ANDREAZA DE MACEDO</sch:nomePagador>
 <sch:textoEnderecoPagador>RUA SEM NOME</sch:textoEnderecoPagador>
 <sch:numeroCepPagador>12345678</sch:numeroCepPagador>
 <sch:nomeMunicipioPagador>BRASILIA</sch:nomeMunicipioPagador>
 <sch:nomeBairroPagador>SIA</sch:nomeBairroPagador>
 <sch:siglaUfPagador>DF</sch:siglaUfPagador>
 <sch:textoNumeroTelefonePagador>45619988</sch:textoNumeroTelefonePagador>
 <sch:codigoTipoInscricaoAvalista/>
 <sch:numeroInscricaoAvalista/>
 <sch:nomeAvalistaTitulo/>
 <sch:codigoChaveUsuario>1</sch:codigoChaveUsuario>
 <sch:codigoTipoCanalSolicitacao>5</sch:codigoTipoCanalSolicitacao>
 </sch:requisicao>
 </soapenv:Body>
</soapenv:Envelope>



******************Resposta do Banco para registrar boletos*********************

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
*/
const registerBoleto = `
 <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
xmlns:sch="http://www.tibco.com/schemas/bws_registro_cbr/Recursos/XSD/Schema.xsd">
 <soapenv:Header/>
 <soapenv:Body>
<sch:requisicao>
 <sch:numeroConvenio>{{.Agreement.AgreementNumber}}</sch:numeroConvenio>
 <sch:numeroCarteira>17</sch:numeroCarteira>
 <sch:numeroVariacaoCarteira>{{.Agreement.WalletVariation}}</sch:numeroVariacaoCarteira>
 <sch:codigoModalidadeTitulo>1</sch:codigoModalidadeTitulo>
 <sch:dataEmissaoTitulo>{{replace (today | brdate) "/" "."}}</sch:dataEmissaoTitulo>
 <sch:dataVencimentoTitulo>{{replace (.Title.ExpireDateTime | brdate) "/" "."}}</sch:dataVencimentoTitulo>
 <sch:valorOriginalTitulo>{{.Title.AmountInCents}}</sch:valorOriginalTitulo>
 <sch:codigoTipoDesconto>0</sch:codigoTipoDesconto> 
 <sch:codigoTipoMulta>0</sch:codigoTipoMulta> 
 <sch:codigoAceiteTitulo>N</sch:codigoAceiteTitulo>
 <sch:codigoTipoTitulo>19</sch:codigoTipoTitulo>
 <sch:textoDescricaoTipoTitulo></sch:textoDescricaoTipoTitulo>
 <sch:indicadorPermissaoRecebimentoParcial>N</sch:indicadorPermissaoRecebimentoParcial>
 <sch:textoNumeroTituloBeneficiario></sch:textoNumeroTituloBeneficiario>
 <sch:textoNumeroTituloCliente>000{{padLeft (toString .Agreement.AgreementNumber) "0" 7}}{{padLeft (toString .Title.OurNumber) "0" 10}}</sch:textoNumeroTituloCliente>
 <sch:textoMensagemBloquetoOcorrencia>Pagamento disponível até a data de vencimento</sch:textoMensagemBloquetoOcorrencia>
 <sch:codigoTipoInscricaoPagador>{{docType .Buyer.Document.Type}}</sch:codigoTipoInscricaoPagador>
 <sch:numeroInscricaoPagador>{{.Buyer.Document.Number}}</sch:numeroInscricaoPagador>
 <sch:nomePagador>{{.Buyer.Name}}</sch:nomePagador>
 <sch:textoEnderecoPagador>{{.Buyer.Address.Street}}</sch:textoEnderecoPagador>
 <sch:numeroCepPagador>{{.Buyer.Address.ZipCode}}</sch:numeroCepPagador>
 <sch:nomeMunicipioPagador>{{.Buyer.Address.City}}</sch:nomeMunicipioPagador>
 <sch:nomeBairroPagador>{{.Buyer.Address.District}}</sch:nomeBairroPagador>
 <sch:siglaUfPagador>{{.Buyer.Address.StateCode}}</sch:siglaUfPagador> 
 <sch:codigoChaveUsuario>1</sch:codigoChaveUsuario>
 <sch:codigoTipoCanalSolicitacao>5</sch:codigoTipoCanalSolicitacao>
 </sch:requisicao>
 </soapenv:Body>
</soapenv:Envelope>
 `

//GetRegisterBoletoBBTmpl retorna o template do Banco do Brasil
func GetRegisterBoletoBBTmpl() string {
	return registerBoleto
}

//GetRegisterBoletoReponseTranslator retorna as regras de tradução da resposta de registrar boleto
func GetRegisterBoletoReponseTranslator() *parser.TranslatorMap {
	translator := parser.NewTranslatorMap()
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:codigoRetornoPrograma", MapKey: "returnCode"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:nomeProgramaErro", MapKey: "errorCode"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:textoMensagemErro", MapKey: "errorMessage"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:numeroPosicaoErroPrograma", MapKey: "positionNumberErrorProgram"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:codigoTipoRetornoPrograma", MapKey: "returnTypeCodeProgram"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:textoNumeroTituloCobrancaBb", MapKey: "numberTextTitleCharging"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:numeroCarteiraCobranca", MapKey: "walletNumberCharging"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:numeroVariacaoCarteiraCobranca", MapKey: "rateNumberWalletCharging"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:codigoPrefixoDependenciaBeneficiario", MapKey: "prefixCodeBeneficiaryDependency"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:numeroContaCorrenteBeneficiario", MapKey: "checkingNumberBeneficiary"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:codigoCliente", MapKey: "clientCode"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:linhaDigitavel", MapKey: "digitableLine"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:codigoBarraNumerico", MapKey: "barcodeNumber"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:codigoTipoEnderecoBeneficiario", MapKey: "addressTypeCodeBeneficiary"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:nomeLogradouroBeneficiario", MapKey: "addressBeneficiary"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:nomeBairroBeneficiario", MapKey: "beneficiaryNeighborhood"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:nomeMunicipioBeneficiario", MapKey: "beneficiaryCity"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:codigoMunicipioBeneficiario", MapKey: "beneficiaryCityCode"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:siglaUfBeneficiario", MapKey: "beneficiaryUfInitials"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:codigoCepBeneficiario", MapKey: "beneficiaryZipCode"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:indicadorComprovacaoBeneficiario", MapKey: "beneficiaryIndicatorEvidence"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:codigoMunicipioBeneficiario", MapKey: "beneficiaryCityCode"})
	translator.AddRule(parser.Rule{XMLQuery: "///ns0:resposta/ns0:numeroContratoCobranca", MapKey: "chargingContractNumber"})
	return translator
}

// GetRegisterBoletoError retorna as regras para ler os campos de erro do banco do brasil
func GetRegisterBoletoError() *parser.TranslatorMap {
	translator := parser.NewTranslatorMap()
	translator.AddRule(parser.Rule{XMLQuery: "//////ns:Mensagem", MapKey: "messageString"})
	translator.AddRule(parser.Rule{XMLQuery: "////faultstring", MapKey: "faultString"})
	translator.AddRule(parser.Rule{XMLQuery: "////faultcode", MapKey: "faultCode"})
	return translator
}

const bbRegisterBoletoResponse = `{
	"beneficiaryZipCode": "{{trim .beneficiaryZipCode}}",
	"numberTextTitleCharging": "{{trim .numberTextTitleCharging}}",
	"rateNumberWalletCharging": "{{trim .rateNumberWalletCharging}}",
	"digitableLine": "{{trim .digitableLine}}",
	"addressTypeCodeBeneficiary": "{{trim .addressTypeCodeBeneficiary}}",
	"walletNumberCharging": "{{trim .walletNumberCharging}}",
	"clientCode": "{{trim .clientCode}}",
	"addressBeneficiary": "{{trim .addressBeneficiary}}",
	"beneficiaryCity": "{{trim .beneficiaryCity}}",
	"prefixCodeBeneficiaryDependency": "{{trim .prefixCodeBeneficiaryDependency}}",
	"checkingNumberBeneficiary": "{{trim .checkingNumberBeneficiary}}",
	"beneficiaryCityCode": "{{trim .beneficiaryCityCode}}",
	"beneficiaryIndicatorEvidence": "{{trim .beneficiaryIndicatorEvidence}}",
	"returnCode": "{{trim .returnCode}}",
	"programError": "{{trim .errorCode}}",
	"positionNumberErrorProgram": "{{trim .positionNumberErrorProgram}}",
	"returnTypeCodeProgram": "{{trim .returnTypeCodeProgram}}",
	"chargingContractNumber": "{{trim .chargingContractNumber}}",
	"errorMessage": "{{trim .errorMessage}}",
	"barcodeNumber": "{{trim .barcodeNumber}}",
	"beneficiaryNeighborhood": "{{trim .beneficiaryNeighborhood}}"
}`

//GetRegisterBoletoBBApiResponseTmpl retorna o template do Banco do Brasil de resposta para a Api
func GetRegisterBoletoBBApiResponseTmpl() string {
	return bbRegisterBoletoResponse
}
