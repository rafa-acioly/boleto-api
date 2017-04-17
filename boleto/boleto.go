package boleto

import (
	"github.com/boombuler/barcode/code128"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/barcode"
	"github.com/jung-kurt/gofpdf/contrib/httpimg"
)

//Create cria um boleto em PDF seguindo os novos padrões
//boleto models.BoletoRequest
func Create(w gin.ResponseWriter) {
	Boleto(w)
}
func drawRows(h, y float64, n float64, pdf *gofpdf.Fpdf) {
	var i float64
	for i = 1; i <= n; i++ {
		pdf.Rect(10, y, 140, h*i, "")
		pdf.Rect(150, y, 50, h*i, "")
	}
}

func drawPaymentLocal(pdf *gofpdf.Fpdf) {
	pdf.MoveTo(10, 27)
	pdf.SetFont("Arial", "B", fontSize-4)
	pdf.Cell(20, 10, "Local de Pagamento")
	pdf.SetFont("Arial", "", fontSize-5.2)
	pdf.MoveTo(10, 34)
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	ht := pdf.PointConvert(fontSize - 5.2)
	str := "ATÉ O VENCIMENTO EM QUALQUER BANCO OU CORRESPONDENTE NÃO BANCÁRIO, APÓS O VENCIMENTO,"
	str += "PAGUE EM QUALQUER BANCO OU CORRESPONDENTE NÃO BANCÁRIO"
	pdf.MultiCell(140, ht, tr(str), "", "J", false)
}

func drawPaymentDate(pdf *gofpdf.Fpdf) {
	pdf.MoveTo(150, 27)
	pdf.SetFont("Arial", "B", fontSize-4)
	pdf.Cell(20, 10, "Data de Pagamento")
	pdf.SetFont("Arial", "", fontSize-4)
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	ht := pdf.PointConvert(fontSize - 4)
	pdf.MoveTo(184, 36)
	pdf.MultiCell(140, ht, tr("01/01/2016"), "", "J", false)
}

const fontSize float64 = 12

func Boleto(w gin.ResponseWriter) {
	pdf := gofpdf.New("P", "mm", "A4", "") // A4 210.0 x 297.0

	pdf.SetFont("Arial", "", fontSize)
	//ht := pdf.PointConvert(fontSize)
	//tr := pdf.UnicodeTranslatorFromDescriptor("") // "" defaults to "cp1252"
	//write := func(str string) {
	//	pdf.MultiCell(150, ht, tr(str), "", "C", false)
	//	pdf.Ln(2 * ht)
	//}
	pdf.SetMargins(15, 20, 15)
	pdf.AddPage()
	//	write("Boleto Banco do Brasil")
	//Desenha as grids

	pdf.Line(60, 30, 60, 22) //linha topo vertical 1
	pdf.Line(80, 30, 80, 22) //linha topo vertical 2
	drawRows(10, 30, 2, pdf)
	drawPaymentLocal(pdf)
	drawPaymentDate(pdf)
	drawRows(25, 70, 2, pdf)

	//Desenha o numero 341-7
	pdf.MoveTo(61, 22)
	pdf.SetFont("Arial", "B", fontSize+6)
	pdf.Cell(20, 10, "341-7") //341-7
	//////////////////////////////////////////

	//Desenha o numero do boleto
	pdf.SetFont("Arial", "B", fontSize-1)
	pdf.MoveTo(83, 21)
	pdf.Cell(20, 10, "34191.12345 67890.101112 13141.516171 8 12345678901112")
	////////////////////////////////////////////////////////////

	pdf.Output(w)
}

func A(w gin.ResponseWriter) {
	pdf := gofpdf.New("P", "mm", "A4", "") // A4 210.0 x 297.0
	var fontSize float64 = 12
	pdf.SetFont("Arial", "", fontSize)
	ht := pdf.PointConvert(fontSize)
	tr := pdf.UnicodeTranslatorFromDescriptor("") // "" defaults to "cp1252"
	write := func(str string) {
		pdf.MultiCell(150, ht, tr(str), "", "C", false)
		pdf.Ln(2 * ht)
	}
	pdf.SetMargins(15, 20, 15)
	pdf.AddPage()
	write("através À noite, vovô Kowalsky vê o ímã cair no pé do pingüim queixoso e vovó" +
		"põe açúcar no chá de tâmaras do jabuti feliz.")

	pdf.Output(w)
}

//Carta na manga
func B(w gin.ResponseWriter) {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetFont("Helvetica", "", 12)
	pdf.SetFillColor(200, 200, 220)
	pdf.AddPage()

	url := "https://github.com/jung-kurt/gofpdf/raw/master/image/logo_gofpdf.jpg?raw=true"
	httpimg.Register(pdf, url, "")
	pdf.Image(url, 15, 15, 267, 0, false, "", 0, "")
	pdf.Output(w)

}

func createPdf() (pdf *gofpdf.Fpdf) {
	pdf = gofpdf.New("L", "mm", "A4", "")
	pdf.SetFont("Helvetica", "", 12)
	pdf.SetFillColor(200, 200, 220)
	pdf.AddPage()
	return
}

func BarCode(w gin.ResponseWriter) {
	pdf := createPdf()
	code := "12345678909876543212345678909876543212345678909"
	bcode, err := code128.Encode(code)

	if err == nil {
		key := barcode.Register(bcode)
		barcode.Barcode(pdf, key, 15, 15, 240, 20, false)
	}
	pdf.Ln(2 * 10)
	pdf.Cell(90, 20, code)
	pdf.Output(w)
}
