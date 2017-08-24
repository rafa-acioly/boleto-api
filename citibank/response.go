package citibank

const registerBoletoResponseCiti = `{
    {{if (hasErrorTags . "errorCode" "errorMessage" "exception")}}
        {{if (hasErrorTags . "exception")}}
            "errors": [
                {                    
                    "code": "{{trim .returnCode}}",
                    "message": "{{trim .returnMessage}}"
                }
            ]
        {{else}}
            "errors": [
                {                    
                    "code": "{{trim .errorCode}}",
                    "message": "{{trim .errorMessage}}"
                }
            ]
        {{end}}
        
    {{else}}
        {{if .digitableLine}}
            "digitable_line": "{{fmtDigitableLine (trim .digitableLine)}}",
        {{end}}
        {{if .barcodeNumber}}
            "barcode_number": "{{trim .barcodeNumber}}"
        {{end}}
        
    {{end}}
}
`

func getAPIResponseCiti() string {
	return registerBoletoResponseCiti
}
