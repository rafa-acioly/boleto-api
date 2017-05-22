package letters

const responseCiti = `
<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope 
    xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
    <soapenv:Body>
        <manutencaocobrancabancaria:SERVICO_SAIDA 
            xmlns:manutencaocobrancabancaria="http://caixa.gov.br/sibar/manutencao_cobranca_bancaria/boleto/externo" 
            xmlns:sibar_base="http://caixa.gov.br/sibar">
            <sibar_base:HEADER>
                <VERSAO>1.0</VERSAO>
                <OPERACAO>{{operation}}</OPERACAO>
                <DATA_HORA>{{datetime}}</DATA_HORA>
            </sibar_base:HEADER>
            <COD_RETORNO>{{returnCode}}</COD_RETORNO>
            <MSG_RETORNO>{{returnMessage}}</MSG_RETORNO>
            <DADOS>
                <EXCECAO>{{exception}}</EXCECAO>
                <CODIGO_BARRAS>{{barcodeNumber}}</CODIGO_BARRAS>
                <LINHA_DIGITAVEL>{{digitableLine}}</LINHA_DIGITAVEL>
                <NOSSO_NUMERO>{{ourNumber}}</NOSSO_NUMERO>
                <URL>{{url}}</URL>
            </DADOS>
        </manutencaocobrancabancaria:SERVICO_SAIDA>
    </soapenv:Body>
</soapenv:Envelope>
`

const registerBoletoCiti = `

## SOAPAction:IncluiBoleto
## Content-Type:text/xml

<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ext="http://caixa.gov.br/sibar/manutencao_cobranca_bancaria/boleto/externo" xmlns:sib="http://caixa.gov.br/sibar">
   <soapenv:Header/>
   <soapenv:Body>
      <ext:SERVICO_ENTRADA>
         <sib:HEADER>
            <VERSAO>1.0</VERSAO>
            <!--Optional:-->
            <AUTENTICACAO>{{.Authentication.AuthorizationToken}}</AUTENTICACAO>
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

//GetRegisterBoletoCitiTmpl retorna o padrão de registro de boleto do Citibank
func GetRegisterBoletoCitiTmpl() string {
	return registerBoletoCiti
}

//GetResponseTemplateCiti retorna o template de mensagem do Citibank
func GetResponseTemplateCiti() string {
	return responseCiti
}
