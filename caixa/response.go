package caixa

//Response focado sna integracao com a Caixa
const registerBoletoResponseCaixa = `{
    {{if (eq .returnCode "1")}}
        "errors":[{
            "code":"{{trim .returnCode}}",
            "message":"{{trim .returnMessage}}"
        }]
    {{else}}
        "digitable_line": "{{fmtDigitableLine (trim .digitableLine)}}",
        "barcode_number": "{{trim .barcodeNumber}}",
        "links": [{
            "href":"{{trim .url}}",
            "rel": "pdf",
            "method":"GET"
        }],        
        "ournumber": "{{trim .ourNumber}}"
    {{end}}
}
`

func getAPIResponseCaixa() string {
	return registerBoletoResponseCaixa
}
