package boleto

import (
	"html/template"

	"bitbucket.org/mundipagg/boletoapi/models"
	"bitbucket.org/mundipagg/boletoapi/tmpl"
	gin "gopkg.in/gin-gonic/gin.v1"
)

const templateBoleto = `
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<style>
    body {
        font-family: "Arial";
		background-color: #fff;
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

    .bankLogo {
        width: 28%;
    }
    
    .boletoNumber {
        width: 62%;
        font-weight: bold;
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
        font-size: 12px;
        font-weight: bold;
    }
    
    .text {
        font-size: 12px;
    }
    
    p.content {
        padding: 0px;
        width: 100%;
        margin: 0px;
        font-size: 12px;
    }
    
    .sideBorders {
        border-left: 1px solid black;
        border-right: 1px solid black;
    }

    hr {
        size: 1;
        border: 1px dashed;
		width: 216mm;
    }

    br {
        content: " ";
        display: block;
        margin: 12px 0;
        line-height: 12px;
    }
</style>

<body>
    {{template "boletoForm" .}}
	<hr/>
	{{template "boletoForm" .}}	
	<center><input type="button" onclick="window.print()" value="Imprimir"></center>
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
                <td class="sideBorders center"><span style="font-size:24px;font-weight:bold;">341-7</span></td>
                <td class="boletoNumber center"><span>34191.12345 67890.101112 13141.516171 8 12345678901112</span></td>
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
                    <p class="content right text" style="font-weight:bold;">01/01/2016</p>
                </td>
            </tr>
            <tr>
                <td width="70%" colspan="6">
                    <span class="title">Nome do Beneficiário / CNPJ / CPF / Endereço:</span>
                    <br/>
                    <table border="0" style="border:none">
                        <tr>
                            <td width="60%"><span class="text">Simulação</span></td>
                            <td><span class="text">CNPJ 01.000.000/0001-00</span></td>
                        </tr>
                    </table>
                    <br/>
                    <span class="text">Rua Simulação, 1 - Jardim Simulação - São Paulo - SP - 10000-000</span>
                </td>
                <td width="30%">
                    <span class="title">Agência/Código Beneficiário</span>
                    <br/>
                    <br/>
                    <p class="content right">1234/12345-1</p>
                </td>
            </tr>

            <tr>
                <td width="15%">
                    <span class="title">Data do Documento</span>
                    <br/>
                    <p class="content center">01/07/2015</p>
                </td>
                <td width="17%" colspan="2">
                    <span class="title">Num. do Documento</span>
                    <br/>
                    <p class="content center">1</p>
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
                    <p class="content center">01/07/2015</p>
                </td>
                <td width="30%">
                    <span class="title">Carteira/Nosso Número</span>
                    <br/>
                    <br/>
                    <p class="content right">157/12345678-9</p>
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
                    <p class="content center">157</p>
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
                    <p class="content center">10,00</p>
                </td>
                <td width="30%">
                    <span class="title">(=) Valor do Documento</span>
                    <br/>
                    <br/>
                    <p class="content right">10,00</p>
                </td>
            </tr>
            <tr>
                <td colspan="6" rowspan="4">
                    <span class="title">Instruções de responsabilidade do BENEFICIÁRIO. Qualquer dúvida sobre este boleto contate o beneficiário.</span>
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
                            <td width="60%"><span class="text"><b>Nome do Pagador: </b> Pagador ABC</span></td>
                            <td><span class="text"><b>CNPJ/CPF: </b> 123.121.001-00</span></td>
                        </tr>
                        <tr>
                            <td><span class="text"><b>Endereço: </b> Avenida Brasil 1234 - Jardim Brasil - São Paulo - SP - 9999999-999</span></td>
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
		<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAZYAAAAyCAYAAAB/Av3aAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAABVvSURBVHhejYoBCiQHDMP6/09frxQNwmtnRhCMFf/zlz//4eRWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9V/6I51Y3l+fo5vIcnVze5A7SZ4I9Z9Z/7WDtnf7D5Vdy5vJOsOdWN+nZtt3yTrDnVjerN8+tblrPfXZyeWPPmda9W/u2c4f1/7qD9Cuhdc7Qnd6R8LZLD61fO6d3JKwdrH/rTmido2f6D5cn/Vf+iOdWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9V/6I51Y3l+fo5vIcnVze5A7SZ4I9Z9Z/7WDtnf7D5Vdy5vJOsOdWN+nZtt3yTrDnVjerN8+tblrPfXZyeWPPmda9W/u2c4f1/7qD9Cuhdc7Qnd6R8LZLD61fO6d3JKwdrH/rTmido2f6D5cn/Vf+iOdWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9V/6I51Y3l+fo5vIcnVze5A7SZ4I9Z9Z/7WDtnf7D5Vdy5vJOsOdWN+nZtt3yTrDnVjerN8+tblrPfXZyeWPPmda9W/u2c4f1/7qD9Cuhdc7Qnd6R8LZLD61fO6d3JKwdrH/rTmido2f6D5cn/Vf+iOdWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9V/6I51Y3l+fo5vIcnVze5A7SZ4I9Z9Z/7WDtnf7D5Vdy5vJOsOdWN+nZtt3yTrDnVjerN8+tblrPfXZyeWPPmda9W/u2c4f1/7qD9Cuhdc7Qnd6R8LZLD61fO6d3JKwdrH/rTmido2f6D5cn/Vf+iOdWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9V/6I51Y3l+fo5vIcnVze5A7SZ4I9Z9Z/7WDtnf7D5Vdy5vJOsOdWN+nZtt3yTrDnVjerN8+tblrPfXZyeWPPmda9W/u2c4f1/7qD9Cuhdc7Qnd6R8LZLD61fO6d3JKwdrH/rTmido2f6D5cn/Vf+iOdWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9V/6I51Y3l+fo5vIcnVze5A7SZ4I9Z9Z/7WDtnf7D5Vdy5vJOsOdWN+nZtt3yTrDnVjerN8+tblrPfXZyeWPPmda9W/u2c4f1/7qD9Cuhdc7Qnd6R8LZLD61fO6d3JKwdrH/rTmido2f6D5cn/Vf+iOdWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9V/6I51Y3l+fo5vIcnVze5A7SZ4I9Z9Z/7WDtnf7D5Vdy5vJOsOdWN+nZtt3yTrDnVjerN8+tblrPfXZyeWPPmda9W/u2c4f1/7qD9Cuhdc7Qnd6R8LZLD61fO6d3JKwdrH/rTmido2f6D5cn/Vf+iOdWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9V/6I51Y3l+fo5vIcnVze5A7SZ4I9Z9Z/7WDtnf7D5Vdy5vJOsOdWN+nZtt3yTrDnVjerN8+tblrPfXZyeWPPmda9W/u2c4f1/7qD9Cuhdc7Qnd6R8LZLD61fO6d3JKwdrH/rTmido2f6D5cn/Vf+iOdWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9V/6I51Y3l+fo5vIcnVze5A7SZ4I9Z9Z/7WDtnf7D5Vdy5vJOsOdWN+nZtt3yTrDnVjerN8+tblrPfXZyeWPPmda9W/u2c4f1/7qD9Cuhdc7Qnd6R8LZLD61fO6d3JKwdrH/rTmido2f6D5cn/Vf+iOdWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9V/6I51Y3l+fo5vIcnVze5A7SZ4I9Z9Z/7WDtnf7D5Vdy5vJOsOdWN+nZtt3yTrDnVjerN8+tblrPfXZyeWPPmda9W/u2c4f1/7qD9Cuhdc7Qnd6R8LZLD61fO6d3JKwdrH/rTmido2f6D5cn/Vf+iOdWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9V/6I51Y3l+fo5vIcnVze5A7SZ4I9Z9Z/7WDtnf7D5Vdy5vJOsOdWN+nZtt3yTrDnVjerN8+tblrPfXZyeWPPmda9W/u2c4f1/7qD9Cuhdc7Qnd6R8LZLD61fO6d3JKwdrH/rTmido2f6D5cn/Vf+iOdWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9V/6I51Y3l+fo5vIcnVze5A7SZ4I9Z9Z/7WDtnf7D5Vdy5vJOsOdWN+nZtt3yTrDnVjerN8+tblrPfXZyeWPPmda9W/u2c4f1/7qD9Cuhdc7Qnd6R8LZLD61fO6d3JKwdrH/rTmido2f6D5cn/Vf+iOdWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9V/6I51Y3l+fo5vIcnVze5A7SZ4I9Z9Z/7WDtnf7D5Vdy5vJOsOdWN+nZtt3yTrDnVjerN8+tblrPfXZyeWPPmda9W/u2c4f1/7qD9Cuhdc7Qnd6R8LZLD61fO6d3JKwdrH/rTmido2f6D5cn/Vf+iOdWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9V/6I51Y3l+fo5vIcnVze5A7SZ4I9Z9Z/7WDtnf7D5Vdy5vJOsOdWN+nZtt3yTrDnVjerN8+tblrPfXZyeWPPmda9W/u2c4f1/7qD9Cuhdc7Qnd6R8LZLD61fO6d3JKwdrH/rTmido2f6D5cn/Vf+iOdWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9V/6I51Y3l+fo5vIcnVze5A7SZ4I9Z9Z/7WDtnf7D5Vdy5vJOsOdWN+nZtt3yTrDnVjerN8+tblrPfXZyeWPPmda9W/u2c4f1/7qD9Cuhdc7Qnd6R8LZLD61fO6d3JKwdrH/rTmido2f6D5cn/Vf+iOdWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9V/6I51Y3l+fo5vIcnVze5A7SZ4I9Z9Z/7WDtnf7D5Vdy5vJOsOdWN+nZtt3yTrDnVjerN8+tblrPfXZyeWPPmda9W/u2c4f1/7qD9Cuhdc7Qnd6R8LZLD61fO6d3JKwdrH/rTmido2f6D5cn/Vf+iOdWN5fn6ObyHJ1c3uQO0meCPWfWf+1g7Z3+w+VXcubyTrDnVjfp2bbd8k6w51Y3qzfPrW5az312cnljz5nWvVv7tnOH9f+6g/QroXXO0J3ekfC2Sw+tXzundySsHax/605onaNn+g+XJ/1X/ojnVjeX5+jm8hydXN7kDtJngj1n1n/tYO2d/sPlV3Lm8k6w51Y36dm23fJOsOdWN6s3z61uWs99dnJ5Y8+Z1r1b+7Zzh/X/uoP0K6F1ztCd3pHwtksPrV87p3ckrB2sf+tOaJ2jZ/oPlyf9/z///PkXZb/t1fffG7EAAAAASUVORK5CYII=" alt="">
		<br/>		
		</div>
    </div>
		
	{{end}}
`

func HTML(w gin.ResponseWriter, boleto models.BoletoRequest) {
	b := tmpl.New()
	type Teste struct {
		BankLogo template.HTML
		Boleto   models.BoletoRequest
	}
	m := Teste{
		BankLogo: template.HTML(logoBB),
		Boleto:   boleto,
	}

	s, _ := b.From(m).To(templateBoleto).Transform(boletoForm)
	w.WriteString(s)
}
