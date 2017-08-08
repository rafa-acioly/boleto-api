package citibank

import (
	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/util"
	"strconv"
	"time"
)

type barcodeNumber struct{
	BankNumber string
	Currency string
	CodProduct string
	DateDueFactor string
	Value string
	OurNumberWithDigit string
	BankAccount string
	Wallet string
}

func casting(boleto *models.BoletoRequest) barcodeNumber {
	Bn := barcodeNumber{}
	date, _ := time.Parse("2006-01-02", boleto.Title.ExpireDate)

	Bn.BankNumber = strconv.Itoa(int(boleto.BankNumber))
	Bn.Currency = "9"
	Bn.CodProduct = "3"
	Bn.DateDueFactor = dateDueFactor(date)
	Bn.Value = util.PadLeft(strconv.Itoa(int(boleto.Title.AmountInCents)), "0",10)
	Bn.OurNumberWithDigit = strconv.Itoa(int(boleto.Title.OurNumber)) + mod11(strconv.Itoa(int(boleto.Title.OurNumber)))
	Bn.OurNumberWithDigit = util.PadLeft(Bn.OurNumberWithDigit, "0",12)
	Bn.BankAccount = boleto.Agreement.Account + boleto.Agreement.AccountDigit
	Bn.Wallet = strconv.Itoa(int(boleto.Agreement.Wallet))

	return Bn
}

func generateBar (boleto *models.BoletoRequest) (string, string) {

	genBar := casting(boleto)

	//Generate CodeBar
	code := genBar.BankNumber + genBar.Currency + genBar.DateDueFactor + genBar.Value + genBar.CodProduct + genBar.Wallet + genBar.BankAccount[1:] + genBar.OurNumberWithDigit
	digitBar := mod11Base9(code)

	//Format CodeBar
	codeb := code[:4] + digitBar + code[4:]

	//Calculates and mount digitable line
	//Group 1
	groupOne := genBar.BankNumber + genBar.Currency + genBar.CodProduct + genBar.Wallet + genBar.BankAccount[:1]
	groupOne = groupOne + mod10(groupOne)
	groupOne = groupOne[:5] + "." + groupOne[5:10]

	//Group 2
	groupTwo := genBar.BankAccount[2:] + genBar.OurNumberWithDigit[:2]
	groupTwo = groupTwo + mod10(groupTwo)
	groupTwo =  groupTwo[:5] + "." + groupTwo[5:11]

	//Group 3
	groupThree := genBar.OurNumberWithDigit[2:]
	groupThree = groupThree + mod10(groupThree)
	groupThree = groupThree[:5] + "." + groupThree[5:11]

	//Group 4
	groupFour := genBar.DateDueFactor + genBar.Value

	//Format digitable line
	digitableLine := groupOne + " " + groupTwo + " " + groupThree + " " + digitBar + " " + groupFour

	return codeb, digitableLine
}









