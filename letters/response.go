package letters

const registerBoletoResponse = `{
    {{if (trim .errorCode) ne ""}}
    "Errors": [
        {
            "Code": "{{trim .errorCode}}",
            "Message": "{{trim .errorMessage}}"
        }
    ]
    {{else}}
    "Url": null,
    "DigitableLine": "{{fmtDigitableLine (trim .digitableLine)}}",
    "BarCodeNumber": "{{trim .barcodeNumber}}"
    {{end}}
}
`

//GetRegisterBoletoAPIResponseTmpl retorna o template de resposta para a Api
func GetRegisterBoletoAPIResponseTmpl() string {
	return registerBoletoResponse
}
