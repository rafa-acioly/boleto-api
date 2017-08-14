package bradesco


const registerBoleto = `
{
    "merchant_id": "{{.Authentication.Username}}",
    "meio_pagamento": "300",
    "pedido": {
        "numero": "{{.Title.DocumentNumber}}",
        "valor": {{.Title.AmountInCents}},
        "descricao": ""
    },
    "comprador": {
        "nome": "{{.Buyer.Name}}",
        "documento": "{{.Buyer.Document.Number}}",
        "endereco": {
            "cep": "{{.Buyer.Address.ZipCode}}",
            "logradouro": "{{.Buyer.Address.Street}}",
            "numero": "{{.Buyer.Address.Number}}",
            "complemento": "{{.Buyer.Address.Complement}}",
            "bairro": "{{.Buyer.Address.District}}",
            "cidade": "{{.Buyer.Address.City}}",
            "uf": "{{.Buyer.Address.StateCode}}"
        },
        "ip": "",
        "user_agent": ""
    },
    "boleto": {
        "beneficiario": "{{.Recipient.Name}}",
        "carteira": "{{.Agreement.Wallet}}",
        "nosso_numero": "{{padLeft (toString .Title.OurNumber) "0" 11}}",
        "data_emissao": "{{enDate today "-"}}",
        "data_vencimento": "{{enDate .Title.ExpireDateTime "-"}}",
        "valor_titulo": {{.Title.AmountInCents}},
        "url_logotipo": "",
        "mensagem_cabecalho": "mensagem de cabecalho",
        "tipo_renderizacao": "2",
        "instrucoes": {
            "instrucao_linha_1": "{{.Title.Instructions}}",
            "instrucao_linha_2": "instrucao 02",
            "instrucao_linha_3": "instrucao 03"
        },
        "registro": {
            "agencia_pagador": "00014",
            "razao_conta_pagador": "07050",
            "conta_pagador": "12345679",
            "controle_participante": "Segurança arquivo remessa",
            "aplicar_multa": true,
            "valor_percentual_multa": 0,
            "valor_desconto_bonificacao": 0,
            "debito_automatico": false,
            "rateio_credito": false,
            "endereco_debito_automatico": "1",
            "tipo_ocorrencia": "02",
            "especie_titulo": "01",
            "primeira_instrucao": "00",
            "segunda_instrucao": "00",
            "valor_juros_mora": 0,
            "data_limite_concessao_desconto": "",
            "valor_desconto": 0,
            "valor_iof": 0,
            "valor_abatimento": 0,
            {{if (eq .Buyer.Document.Type "CPF")}}
            	"tipo_inscricao_pagador": "02",
			{{else}}
            	"tipo_inscricao_pagador": "01",
			{{end}}
            "sequencia_registro": "00001"
        }
    },
    "token_request_confirmacao_pagamento": "21323dsd23434ad12178DDasY"
}
`

const responseBradesco = `
{
    "merchant_id": "",
    "meio_pagamento": "",
    "pedido": {
        "numero": "",
        "valor": 0,
        "descricao": ""
    },
    "boleto": {
        "valor_titulo": 0,
        "data_geracao": "{{datetime}}",
        "data_atualizacao": null,
        "linha_digitavel": "",
        "linha_digitavel_formatada": "{{digitableLine}}",
        "token": "",
        "url_acesso": "{{url}}"
    },
    "status": {
        "codigo": {{returnCode}},
        "mensagem": "{{returnMessage}}",
        "detalhes": null
    }
}




`