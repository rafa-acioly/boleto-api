package letters

const registerBoletoResponse = `{
    {{if (hasErrorTags . "errorCode" "errorMessage" "exception")}}
        {{if (hasErrorTags . "exception")}}
            "Errors": [
                {                    
                    "Code": "{{trim .returnCode}}",
                    "Message": "{{trim .returnMessage}}"
                }
            ]
        {{else}}
            "Errors": [
                {                    
                    "Code": "{{trim .errorCode}}",
                    "Message": "{{trim .errorMessage}}"
                }
            ]
        {{end}}
        
    {{else}}
        {{if .digitableLine}}
            "DigitableLine": "{{fmtDigitableLine (trim .digitableLine)}}",
        {{end}}
        {{if .barcodeNumber}}
            "BarCodeNumber": "{{trim .barcodeNumber}}"
        {{end}}
        {{if .url}}
            ,"Links": [{
                "href":"{{trim .url}}",
                "rel": "pdf",
                "method":"GET"
            }], 
        {{end}}
        {{if .ourNumber}}
            "OurNumber": "{{trim .ourNumber}}"
        {{end}}
        
    {{end}}
}
`

//GetRegisterBoletoAPIResponseTmpl retorna o template de resposta para a Api
func GetRegisterBoletoAPIResponseTmpl() string {
	return registerBoletoResponse
}
