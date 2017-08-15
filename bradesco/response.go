package bradesco

var apiResponse = `
{

//É preciso configurar essa parte de erro e testar
	{{if (hasErrorTags . "errorCode")}}
        "Errors": [
            {
                "Code": "{{trim .errorCode}}",
                "Message": "{{trim .errorMessage}}"
            }
        ]
    {{else}}
		"DigitableLine": "{{.digitableLine}}",
		"Links": [{
			"href":"{{trim .url}}",
			"rel": "pdf",
			"method":"GET"
		}]
    {{end}}
}
`

func getAPIResponseBradesco() string {
	return apiResponse
}
