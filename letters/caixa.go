package letters

const incluiBoleto = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ext="http://caixa.gov.br/sibar/manutencao_cobranca_bancaria/boleto/externo" xmlns:sib="http://caixa.gov.br/sibar">
   <soapenv:Header/>
   <soapenv:Body>
      <ext:SERVICO_ENTRADA>
         <sib:HEADER>
            <VERSAO>?</VERSAO>
            <AUTENTICACAO>0200400000000000000000000510201600000000015005000123456789124</AUTENTICACAO>            
            <SISTEMA_ORIGEM>SIGCB</SISTEMA_ORIGEM>           
         </sib:HEADER>
         <DADOS>
            <!--You have a CHOICE of the next 3 items at this level-->
            <INCLUI_BOLETO>
               <CODIGO_BENEFICIARIO>{{.Agreement.AgreementNumber}}</CODIGO_BENEFICIARIO>
               <TITULO>                  
                  <NOSSO_NUMERO>{{.Title.OurNumber}}</NOSSO_NUMERO>
                  <NUMERO_DOCUMENTO>{{.Title.DocumentNumber}}</NUMERO_DOCUMENTO>
                  <DATA_VENCIMENTO>{{brDateWithoutDelimiter .Title.ExpireDateTime}}</DATA_VENCIMENTO>
                  <VALOR>{{.Title.AmountInCents}}</VALOR>
                  <TIPO_ESPECIE>17</TIPO_ESPECIE>
                  <FLAG_ACEITE>S</FLAG_ACEITE>                                                                       
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

const baixaBoleto = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ext="http://caixa.gov.br/sibar/manutencao_cobranca_bancaria/boleto/externo" xmlns:sib="http://caixa.gov.br/sibar">
   <soapenv:Header/>
   <soapenv:Body>
      <ext:SERVICO_ENTRADA>
         <sib:HEADER>
            <VERSAO>?</VERSAO>
            <AUTENTICACAO>0200400000000000000000000510201600000000015005000123456789124</AUTENTICACAO>            
            <SISTEMA_ORIGEM>SIGCB</SISTEMA_ORIGEM>           
         </sib:HEADER>
         <DADOS>            
            <BAIXA_BOLETO>
               <CODIGO_BENEFICIARIO>{{.Agreement.AgreementNumber}}</CODIGO_BENEFICIARIO>
               <NOSSO_NUMERO>{{.Title.OurNumber}}</NOSSO_NUMERO>
            </BAIXA_BOLETO>            
         </DADOS>
      </ext:SERVICO_ENTRADA>
   </soapenv:Body>
</soapenv:Envelope>
`

const alteraBoleto = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ext="http://caixa.gov.br/sibar/manutencao_cobranca_bancaria/boleto/externo" xmlns:sib="http://caixa.gov.br/sibar">
   <soapenv:Header/>
   <soapenv:Body>
      <ext:SERVICO_ENTRADA>
         <sib:HEADER>
            <VERSAO>?</VERSAO>
            <AUTENTICACAO>0200400000000000000000000510201600000000015005000123456789124</AUTENTICACAO>            
            <SISTEMA_ORIGEM>SIGCB</SISTEMA_ORIGEM>           
         </sib:HEADER>
         <DADOS>            
            <ALTERA_BOLETO>
               <CODIGO_BENEFICIARIO>?</CODIGO_BENEFICIARIO>
               <TITULO>                  
                  <NOSSO_NUMERO>{{.Title.OurNumber}}</NOSSO_NUMERO>
                  <NUMERO_DOCUMENTO>{{.Title.DocumentNumber}}</NUMERO_DOCUMENTO>
                  <DATA_VENCIMENTO>{{.Title.ExpireDate}}</DATA_VENCIMENTO>
                  <VALOR>{{.Title.AmountInCents}}</VALOR>
                  <TIPO_ESPECIE>17</TIPO_ESPECIE>
                  <FLAG_ACEITE>S</FLAG_ACEITE>                                                                       
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
            </ALTERA_BOLETO>
         </DADOS>
      </ext:SERVICO_ENTRADA>
   </soapenv:Body>
</soapenv:Envelope>
`

//GetRegisterBoletoCaixaTmpl retorna o padrão de registro de boleto da Caixa
func GetRegisterBoletoCaixaTmpl() string {
	return incluiBoleto
}
