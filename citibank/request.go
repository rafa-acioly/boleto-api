package citibank

const responseCiti = `
<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:s1="http://www.citibank.com.br/comercioeletronico/registerboleto">
    <soap:Body>
        <s1:RegisterBoletoResponse>
            <actionCode>{{returnCode}}</actionCode>
            <reasonMessage>{{returnMessage}}</reasonMessage>
            <TitlBarCd>{{barcodeNumber}}</TitlBarCd>
            <TitlDgtLine>{{digitableLine}}</TitlDgtLine>
        </s1:RegisterBoletoResponse>
    </soap:Body>
</soap:Envelope>
`

const registerBoletoCiti = `

## SOAPAction:RegisterBoleto
## Content-Type:text/xml

<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
   <soapenv:Header/>
   <soapenv:Body>
      <GrpREMColTit>
         <GrpBenf>
            <CdClrSys>745</CdClrSys>
            <CdIspb>33479023</CdIspb>
            <CdtrId>{{.Authentication.Username}}</CdtrId>
            <CdtrNm>{{.Recipient.Name}}</CdtrNm>
            <CdtrTaxId>{{.Recipient.Document.Number}}</CdtrTaxId>
            <CdtrTaxTp>J</CdtrTaxTp>
         </GrpBenf>
         <GrpClPgd>
            <DbtrNm>{{.Buyer.Name}}</DbtrNm>
            <DbtrTaxId>{{.Buyer.Document.Number}}</DbtrTaxId>
			{{if (eq .Buyer.Document.Type "CPF")}}
            	<DbtrTaxTp>F</DbtrTaxTp>
			{{else}}
            	<DbtrTaxTp>J</DbtrTaxTp>
			{{end}}
            <GrpClPgdAdr>
               <DbtrAdrTp>{{.Buyer.Address.Street}} {{.Buyer.Address.Number}} {{.Buyer.Address.Complement}}</DbtrAdrTp>
               <DbtrCtrySubDvsn>{{.Buyer.Address.StateCode}}</DbtrCtrySubDvsn>
               <DbtrPstCd>{{.Buyer.Address.ZipCode}}</DbtrPstCd>
               <DbtrTwnNm>{{.Buyer.Address.City}}</DbtrTwnNm>
            </GrpClPgdAdr>
         </GrpClPgd>
         <CdOccTp>01</CdOccTp>
         <DbtrGrntNm> </DbtrGrntNm>
         <DbtrMsg>{{.Title.Instructions}}</DbtrMsg>
         <TitlAmt>{{.Title.AmountInCents}}</TitlAmt>
         <TitlBarCdInd>0</TitlBarCdInd>
         <TitlCcyCd>09</TitlCcyCd>
         <TitlCiaCdId>{{trimLeft .Title.DocumentNumber "0"}}</TitlCiaCdId>
         <TitlDueDt>{{enDate .Title.ExpireDateTime "-"}}</TitlDueDt>
         <TitlInstrNmDtExec>0</TitlInstrNmDtExec>
         <TitlInstrProtInd> </TitlInstrProtInd>
         <TitlInstrWrtOffInd> </TitlInstrWrtOffInd>
         <TitlIOFAmt>0</TitlIOFAmt>
         <TitlIssDt>{{enDate todayCiti "-"}}</TitlIssDt>
         <TitlOurNb>{{padLeft (toString .Title.OurNumber) "0" 12}}</TitlOurNb>
         <TitlPortCd>1</TitlPortCd>
         <TitlRbtAmt>0</TitlRbtAmt>
         <TitlTpCd>03</TitlTpCd>
         <TitlYourNb>{{trimLeft .Title.DocumentNumber "0"}}</TitlYourNb>
         <GrpDscnt>
            <TitlDscntAmtOrPrct>0</TitlDscntAmtOrPrct>
            <TitlDscntEndDt> </TitlDscntEndDt>
            <TitlDscntTp> </TitlDscntTp>
         </GrpDscnt>
         <GrpItrs>
            <TitlItrsAmtOrPrct>0</TitlItrsAmtOrPrct>
            <TitlItrsStrDt> </TitlItrsStrDt>
            <TitlItrsTp> </TitlItrsTp>
         </GrpItrs>
         <GrpFn>
            <TitlFnAmtOrPrct>0</TitlFnAmtOrPrct>
            <TitlFnStrDt> </TitlFnStrDt>
            <TitlFnTp> </TitlFnTp>
         </GrpFn>
      </GrpREMColTit>
   </soapenv:Body>
</soapenv:Envelope>
`

func getRequestCiti() string {
	return registerBoletoCiti
}

func getResponseCiti() string {
	return responseCiti
}
