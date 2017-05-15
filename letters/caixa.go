package letters

const erro = `
<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope 
    xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
    <soapenv:Body>
        <manutencaocobrancabancaria:SERVICO_SAIDA 
            xmlns:manutencaocobrancabancaria="http://caixa.gov.br/sibar/manutencao_cobranca_bancaria/boleto/externo" 
            xmlns:sibar_base="http://caixa.gov.br/sibar">
            <sibar_base:HEADER>
                <VERSAO>1.0</VERSAO>
                <OPERACAO>SEM_OPERACAO</OPERACAO>
                <DATA_HORA>20170515125658</DATA_HORA>
            </sibar_base:HEADER>
            <COD_RETORNO>X5</COD_RETORNO>
            <ORIGEM_RETORNO>BROKER-SPD7CJD0</ORIGEM_RETORNO>
            <MSG_RETORNO>(BK76) ERRO NA FORMATACAO DA MENSAGEM.</MSG_RETORNO>
            <DADOS>
                <EXCECAO>EXCECAO NO BAR_MANUTENCAO_COBRANCA_BANCARIA_WS.SOAPInput_Empresas_Externas. DETALHES: ParserException(1) - Funcao: ImbDataFlowNode::createExceptionList, Texto Excecao: Node throwing exception, Texto de Insercao(1) - BAR_MANUTENCAO_COBRANCA_BANCARIA_WS.SOAPInput_Empresas_Externas.ParserException(2) - Funcao: ImbSOAPInputNode::validateData, Texto Excecao: Error occurred in ImbSOAPInputHelper::validateSOAPInput().ParserException(3) - Funcao: ImbRootParser::parseNextItem, Texto Excecao: Exception whilst parsing.ParserException(4) - Funcao: ImbSOAPParser::createSoapShapedTree, Texto Excecao: problem creating SOAP tree from bitstream.ParserException(5) - Funcao: ImbXMLNSCParser::parseLastChild, Texto Excecao: XML Parsing Errors have occurred.ParserException(6) - Funcao: ImbXMLNSCDocHandler::handleParseErrors, Texto Excecao: A schema validation error has occurred while parsing the XML document, Texto de Insercao(1) - 5008, Texto de Insercao(2) - 2, Texto de Insercao(3) - 9, Texto de Insercao(4) - 29, Texto de Insercao(5) - cvc-complex-type.2.4.e: Unexpected element. Element "SISTEMA_ORIGEM" is not one of the choices., Texto de Insercao(6) - /XMLNSC/http://schemas.xmlsoap.org/soap/envelope/:Envelope/http://schemas.xmlsoap.org/soap/envelope/:Body/http://caixa.gov.br/sibar/manutencao_cobranca_bancaria/boleto/externo:SERVICO_ENTRADA/http://caixa.gov.br/sibar:HEADER.</EXCECAO>
            </DADOS>
        </manutencaocobrancabancaria:SERVICO_SAIDA>
    </soapenv:Body>
</soapenv:Envelope>
`

const incluiBoleto = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ext="http://caixa.gov.br/sibar/manutencao_cobranca_bancaria/boleto/externo" xmlns:sib="http://caixa.gov.br/sibar">
   <soapenv:Header/>
   <soapenv:Body>
      <ext:SERVICO_ENTRADA>
         <sib:HEADER>
            <VERSAO>1.0</VERSAO>
            <!--Optional:-->
            <AUTENTICACAO>{{.Authentication.Password}}</AUTENTICACAO>
            <USUARIO>1234567</USUARIO>
            <OPERACAO>INCLUI_BOLETO</OPERACAO>
            <SISTEMA_ORIGEM>SIGCB</SISTEMA_ORIGEM>
            <DATA_HORA>{{fullDate today}}</DATA_HORA>
         </sib:HEADER>
         <DADOS>
            <!--You have a CHOICE of the next 3 items at this level-->
            <INCLUI_BOLETO>
               <CODIGO_BENEFICIARIO>{{.Agreement.AgreementNumber}}</CODIGO_BENEFICIARIO>
               <TITULO>                  
                  <NOSSO_NUMERO>{{.Title.OurNumber}}</NOSSO_NUMERO>
                  <NUMERO_DOCUMENTO>{{.Title.DocumentNumber}}</NUMERO_DOCUMENTO>
                  <DATA_VENCIMENTO>{{enDate .Title.ExpireDateTime "-"}}</DATA_VENCIMENTO>
                  <VALOR>{{.Title.AmountInCents}}</VALOR>
                  <TIPO_ESPECIE>17</TIPO_ESPECIE>
                  <FLAG_ACEITE>S</FLAG_ACEITE>   
                  <JUROS_MORA>
                     <TIPO>ISENTO</TIPO>
                     <VALOR>0</VALOR>                                                        
                  </JUROS_MORA>
                  <POS_VENCIMENTO>
                     <ACAO>DEVOLVER</ACAO>
                     <NUMERO_DIAS>30</NUMERO_DIAS>
                  </POS_VENCIMENTO>                       
                  <CODIGO_MOEDA>9</CODIGO_MOEDA>
                  <PAGADOR>
                     <!--You have a CHOICE of the next 2 items at this level-->
                     {{if eq .Buyer.Document.Type "CPF"}}
					 	<CPF>{{.Buyer.Document.Number}}</CPF>
                     	<NOME>{{.Buyer.Name}}</NOME>
                     {{else}}
					 	<CNPJ>{{.Buyer.Document.Number}}</CNPJ>
                     	<RAZAO_SOCIAL>{{.Buyer.Name}}</RAZAO_SOCIAL>
					 {{end}}
                     <ENDERECO>
                        <LOGRADOURO>{{.Buyer.Address.Street}} {{.Buyer.Address.Number}} {{.Buyer.Address.Complement}}</LOGRADOURO>
                        <BAIRRO>{{.Buyer.Address.District}}</BAIRRO>
                        <CIDADE>{{.Buyer.Address.City}}</CIDADE>
                        <UF>{{.Buyer.Address.StateCode}}</UF>
                        <CEP>{{.Buyer.Address.ZipCode}}</CEP>
                     </ENDERECO>
                  </PAGADOR>                                                                       
                  <FICHA_COMPENSACAO>
                     <MENSAGENS>
                        <MENSAGEM>{{.Title.Instructions}}</MENSAGEM>
                     </MENSAGENS>
                  </FICHA_COMPENSACAO>                  
               </TITULO>
            </INCLUI_BOLETO>
         </DADOS>
      </ext:SERVICO_ENTRADA>
   </soapenv:Body>
</soapenv:Envelope>
`

//GetRegisterBoletoCaixaTmpl retorna o padrão de registro de boleto da Caixa
func GetRegisterBoletoCaixaTmpl() string {
	return incluiBoleto
}
