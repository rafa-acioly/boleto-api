package bb

const registerBoletoResponseBB = `{
    {{if (hasErrorTags . "errorCode")}}
        "errors": [
            {                    
                "code": "{{trim .errorCode}}",
                "message": "{{trim .errorMessage}}"
            }
        ]
    {{else}}
        "digitable_line": "{{fmtDigitableLine (trim .digitableLine)}}",
        "barcode_number": "{{trim .barcodeNumber}}"
    {{end}}
}
`

func getAPIResponse() string {
	return registerBoletoResponseBB
}
