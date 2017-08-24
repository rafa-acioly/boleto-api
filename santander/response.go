package santander

var apiResponse = `
{
    {{if (hasErrorTags . "errorCode")}}
        "errors": [
            {                    
                "code": "{{trim .errorCode}}",
                "message": "{{trim .errorMessage}}"
            }
        ]
    {{else}}
        "digitable_line": "{{fmtDigitableLine (trim .digitableLine)}}",
        "barcode_number": "{{trim .barcodeNumber}}",
        "ournumber":"{{.ourNumber}}"
    {{end}}
}
`

func getAPIResponseSantander() string {
	return apiResponse
}
