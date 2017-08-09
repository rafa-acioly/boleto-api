package santander

var apiResponse = `
{
    {{if (hasErrorTags . "errorCode")}}
        "Errors": [
            {                    
                "Code": "{{trim .errorCode}}",
                "Message": "{{trim .errorMessage}}"
            }
        ]
    {{else}}
        "DigitableLine": "{{fmtDigitableLine (trim .digitableLine)}}",
        "BarCodeNumber": "{{trim .barcodeNumber}}",
        "OurNumber":"{{.ourNumber}}"
    {{end}}
}
`

func getAPIResponseSantander() string {
	return apiResponse
}
