package letters

const registerBoletoResponse = `{
    "Url": null,
    "DigitableLine": "{{trim .digitableLine}}",
    "BarCodeNumber": "{{trim .barcodeNumber}}",
    "Errors": [
        {
            "Code": "{{trim .errorCode}}",
            "Message": "{{trim .errorMessage}}"
        }
    ]
}
`

//GetRegisterBoletoAPIResponseTmpl retorna o template de resposta para a Api
func GetRegisterBoletoAPIResponseTmpl() string {
	return registerBoletoResponse
}
