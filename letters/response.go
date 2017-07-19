package letters

import "github.com/mundipagg/boleto-api/models"

const registerBoletoResponseBB = `{
    {{if (hasErrorTags . "errorCode")}}
        "Errors": [
            {                    
                "Code": "{{trim .errorCode}}",
                "Message": "{{trim .errorMessage}}"
            }
        ]
    {{else}}
        "DigitableLine": "{{fmtDigitableLine (trim .digitableLine)}}",
        "BarCodeNumber": "{{trim .barcodeNumber}}"
    {{end}}
}
`

const registerBoletoResponseCiti = `{
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
        
    {{end}}
}
`

//Response focado sna integracao com a Caixa
const registerBoletoResponseCaixa = `{
    {{if (eq .returnCode "1")}}
        "Errors":[{
            "Code":"{{trim .returnCode}}",
            "Message":"{{trim .returnMessage}}"
        }]
    {{else}}
        "DigitableLine": "{{fmtDigitableLine (trim .digitableLine)}}",
        "BarCodeNumber": "{{trim .barcodeNumber}}",
        "Links": [{
            "href":"{{trim .url}}",
            "rel": "pdf",
            "method":"GET"
        }],        
        "OurNumber": "{{trim .ourNumber}}"
    {{end}}
}
`

//GetRegisterBoletoAPIResponseTmpl retorna o template de resposta para a Api
func GetRegisterBoletoAPIResponseTmpl(bank models.BankNumber) string {
	switch bank {
	case models.Caixa:
		return registerBoletoResponseCaixa
	case models.BancoDoBrasil:
		return registerBoletoResponseBB
	case models.Citibank:
		return registerBoletoResponseCiti
	default:
		return ""
	}
}
