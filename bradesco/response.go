package bradesco

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
        "Links": [{
            "href":"{{trim .url}}",
            "rel": "pdf",
            "method":"GET"
        }],
    {{end}}
}
`

func getAPIResponseBradesco() string {
	return apiResponse
}
