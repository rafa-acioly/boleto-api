package boleto

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/twooffive"

	"image/jpeg"

	"bitbucket.org/mundipagg/boletoapi/models"
	"bitbucket.org/mundipagg/boletoapi/tmpl"
)

const templateBoleto = `
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <style>

        @media print
        {    
            .no-print, .no-print *
            {
                display: none !important;
            }
        }

        body {
            font-family: "Arial";
    		background-color: #fff;
            font-size:0.7em;
        }
        .left {
    		margin: auto;		
    		width: 216mm;
    	}
        .document {
            margin: auto auto;
            width: 216mm;
            height: 108mm;
        }

        table {
            width: 100%;
            position: relative;
            border-collapse: collapse;
        }

        .boletoNumber {
            width: 66%;
            font-weight: bold;
            font-size:0.9em;
        }

        .center {
            text-align: center;
        }

        .right {
            text-align: right;
            right: 20px;
        }

        td {
            position: relative;
        }

        .title {
            position: absolute;
            left: 0px;
            top: 0px;
            font-size:0.65em;
            font-weight: bold;
        }

        .text {
             font-size:0.7em;
        }

        p.content {
            padding: 0px;
            width: 100%;
            margin: 0px;
            font-size:0.7em;
        }

        .sideBorders {
            border-left: 1px solid black;
            border-right: 1px solid black;
        }

        hr {
            size: 1;
            border: 1px dashed;
    		width: 216mm;
    		margin-top: 9mm;
        	margin-bottom: 9mm;
        }

        br {
            content: " ";
            display: block;
            margin: 12px 0;
            line-height: 12px;
        }

        .print {
            /* TODO(dbeam): reconcile this with overlay.css' .default-button. */
            background-color: rgb(77, 144, 254);
            background-image: linear-gradient(to bottom, rgb(77, 144, 254), rgb(71, 135, 237));
            border: 1px solid rgb(48, 121, 237);
            color: #fff;
            text-shadow: 0 1px rgba(0, 0, 0, 0.1);
        }

        .btnDefault {
            font-kerning: none;
            font-weight: bold;
        }

        .btnDefault:not(:focus):not(:disabled) {
            border-color: #808080;
        }

        button {
            padding: 5px;
            line-height: 20px;
        }

        span.iconFont {
            font-size: 20px;
        }

        span.align {
            display: inline-block;
            vertical-align: middle;
        }
    </style>
    <link rel="stylesheet" href="http://code.ionicframework.com/ionicons/2.0.1/css/ionicons.min.css">
</head>

<body>
    {{if eq .Format "html"}}	
	    <center><input class="no-print" type="button" onclick="window.print()" value="Imprimir"></center>
    {{end}}
    {{template "boletoForm" .}}

	<hr/>
	{{template "boletoForm" .}}
    </div>	
</body>

</html>
`

const boletoForm = `
{{define "boletoForm"}}
<div class="document">
        <table cellspacing="0" cellpadding="0">
            <tr class="topLine">
                <td class="bankLogo">
                    {{.BankLogo}}					
                </td>
                <td class="sideBorders center"><span style="font-weight:bold;font-size:0.9em;">{{.BankNumber}}</span></td>
                <td class="boletoNumber center"><span>{{fmtDigitableLine .DigitableLine}}</span></td>
            </tr>
        </table>
        <table cellspacing="0" cellpadding="0" border="1">
            <tr>
                <td width="70%" colspan="6">
                    <span class="title">Local de Pagamento</span>
                    <br/>
                    <span class="text">ATÉ O VENCIMENTO EM QUALQUER BANCO OU CORRESPONDENTE NÃO BANCÁRIO, APÓS O VENCIMENTO, PAGUE EM QUALQUER BANCO OU CORRESPONDENTE NÃO BANCÁRIO</span>
                </td>
                <td width="30%">
                    <span class="title">Data de Vencimento</span>
                    <br/>
                    <br/>
                    <p class="content right text" style="font-weight:bold;">{{.Boleto.Title.ExpireDateTime | brdate}}</p>
                </td>
            </tr>
            <tr>
                <td width="70%" colspan="6">
                    <span class="title">Nome do Beneficiário / CNPJ / CPF / Endereço:</span>
                    <br/>
                    <table border="0" style="border:none">
                        <tr>
                            <td width="60%"><span class="text">{{.Boleto.Recipient.Name}}</span></td>
                            <td><span class="text"><b>{{.Boleto.Recipient.Document.Type}}</b> {{fmtDoc .Boleto.Recipient.Document}}</span></td>
                        </tr>
                    </table>
                    <br/>
                    <span class="text">{{.Boleto.Recipient.Address.Street}}, 
                    {{.Boleto.Recipient.Address.Number}} - 
                    {{.Boleto.Recipient.Address.District}}, 
                    {{.Boleto.Recipient.Address.StateCode}} - 
                    {{.Boleto.Recipient.Address.ZipCode}}</span>
                </td>
                <td width="30%">
                    <span class="title">Agência/Código Beneficiário</span>
                    <br/>
                    <br/>
                    <p class="content right">{{.Boleto.Agreement.Agency}}/{{.Boleto.Agreement.Account}}-{{.Boleto.Agreement.AccountDigit}}</p>
                </td>
            </tr>

            <tr>
                <td width="15%">
                    <span class="title">Data do Documento</span>
                    <br/>
                    <p class="content center">{{today | brdate}}</p>
                </td>
                <td width="17%" colspan="2">
                    <span class="title">Num. do Documento</span>
                    <br/>
                    <p class="content center">{{.Boleto.Title.DocumentNumber}}</p>
                </td>
                <td width="10%">
                    <span class="title">Espécie doc</span>
                    <br/>
                    <p class="content center">DM</p>
                </td>
                <td width="8%">
                    <span class="title">Aceite</span>
                    <br/>
                    <p class="content center">N</p>
                </td>
                <td>
                    <span class="title">Data Processamento</span>
                    <br/>
                    <p class="content center">{{today | brdate}}</p>
                </td>
                <td width="30%">
                    <span class="title">Carteira/Nosso Número</span>
                    <br/>
                    <br/>
                    <p class="content right">17/{{.Boleto.Title.OurNumber}}</p>
                </td>
            </tr>

            <tr>
                <td width="15%">
                    <span class="title">Uso do Banco</span>
                    <br/>
                    <p class="content center">&nbsp;</p>
                </td>
                <td width="10%">
                    <span class="title">Carteira</span>
                    <br/>
                    <p class="content center">17</p>
                </td>
                <td width="10%">
                    <span class="title">Espécie</span>
                    <br/>
                    <p class="content center">R$</p>
                </td>
                <td width="8%" colspan="2">
                    <span class="title">Quantidade</span>
                    <br/>
                    <p class="content center">N</p>
                </td>
                <td>
                    <span class="title">Valor</span>
                    <br/>
                    <p class="content center">{{fmtNumber .Boleto.Title.AmountInCents}}</p>
                </td>
                <td width="30%">
                    <span class="title">(=) Valor do Documento</span>
                    <br/>
                    <br/>
                    <p class="content right">{{fmtNumber .Boleto.Title.AmountInCents}}</p>
                </td>
            </tr>
            <tr>
                <td colspan="6" rowspan="4">
                    <span class="title">Instruções de responsabilidade do BENEFICIÁRIO. Qualquer dúvida sobre este boleto contate o beneficiário.</span>
                    <p class="content">{{.Boleto.Title.Instructions}}</p>
                </td>
            </tr>
            <tr>
                <td>
                    <span class="title">(-) Descontos/Abatimento</span>
                    <br/>
                    <p class="content right">&nbsp;</p>
                </td>
            </tr>
            <tr>
                <td>
                    <span class="title">(+) Juros/Multa</span>
                    <br/>
                    <p class="content right">&nbsp;</p>
                </td>
            </tr>
            <tr>
                <td>
                    <span class="title">(=) Valor Pago</span>
                    <br/>
                    <p class="content right">&nbsp;</p>
                </td>
            </tr>
            <tr>
                <td colspan="7">
                    <table border="0" style="border:none">
                        <tr>
                            <td width="60%"><span class="text"><b>Nome do Pagador: </b>&nbsp;{{.Boleto.Buyer.Name}}</span></td>
                            <td><span class="text"><b>CNPJ/CPF: </b>&nbsp;{{fmtDoc .Boleto.Buyer.Document}}</span></td>
                        </tr>
                        <tr>
                            <td><span class="text"><b>Endereço: </b>&nbsp;{{.Boleto.Buyer.Address.Street}}&nbsp;{{.Boleto.Buyer.Address.Number}}, {{.Boleto.Buyer.Address.District}} - {{.Boleto.Buyer.Address.City}}, {{.Boleto.Buyer.Address.StateCode}} - {{.Boleto.Buyer.Address.ZipCode}}</span></td>
                            <td>&nbsp;</td>
                        </tr>
                        <tr>
                            <td><span class="text"><b>Sacador/Avalista: </b> &nbsp;</span></td>
                            <td><span class="text"><b>CNPJ/CPF: </b> &nbsp;</span></td>
                        </tr>
                    </table>

                </td>

            </tr>
        </table>
		<br/>
		<div class="left">
		<img style="margin-left:5mm;" src="data:image/jpg;base64,{{.Barcode64}}" alt="">
		<br/>		
		</div>
    </div>

	{{end}}
`

//HTML renderiza HTML do boleto
func HTML(boleto models.BoletoView, format string) string {
	b := tmpl.New()
	boleto.BankLogo = template.HTML(logoBB)
	boleto.Format = format
	bcode, _ := twooffive.Encode(boleto.Barcode, true)
	orgBounds := bcode.Bounds()
	orgWidth := orgBounds.Max.X - orgBounds.Min.X
	img, _ := barcode.Scale(bcode, orgWidth, 50)
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, img, nil)
	boleto.Barcode64 = base64.StdEncoding.EncodeToString(buf.Bytes())
	s, err := b.From(boleto).To(templateBoleto).Transform(boletoForm)
	if err != nil {
		fmt.Println(err)
	}
	return s
}
