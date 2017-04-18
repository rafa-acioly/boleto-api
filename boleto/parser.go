package boleto

import (
	"fmt"

	"github.com/beevik/etree"
	"github.com/jung-kurt/gofpdf"
)

const example string = `
	<page leftmargin="20" topmargin="20" rightmargin="20" font-size="12" font="Arial">
		<table x="0" y="0">
			<row height="10" width="100%">
				<col width="30%">
					<label style="bold" x="0" y="5">Local de Pagamento</label>
					<br/>
					<label>ATÈ O VENCIMENTO....</label>
				</col>
				<col width="33%">
					<label style="bold" x="0" y="5">Data de Pagamento</label>
					<br/>
					<label font-size="-3">01/01/2016</label>
				</col>
			</row>
		</table>
	</page>
	
`

var nodeMap = map[string]func(*etree.Element){
	"page":    page,
	"table":   table,
	"row":     row,
	"col":     col,
	"br":      br,
	"label":   label,
	"barcode": _barcode,
}

type RenderEngine struct {
	pdf *gofpdf.Fpdf
}

func (r *RenderEngine) Visitor(child *etree.Element) {
	nodeMap[child.Tag](child)
}

func Parse() {
	engine := RenderEngine{}
	doc := etree.NewDocument()
	if err := doc.ReadFromString(example); err != nil {
		panic(err)
	}
	for _, chd := range doc.ChildElements() {
		walker(chd, engine.Visitor)
	}
}

func walker(child *etree.Element, process func(*etree.Element)) {
	process(child)
	if len(child.ChildElements()) == 0 {
		return
	}
	for _, chd := range child.ChildElements() {
		walker(chd, process)
	}
}

func page(child *etree.Element) {
	fmt.Println("Processa Pagina" + child.Tag)
}

func table(child *etree.Element) {
	fmt.Println("Processa Tabela" + child.Tag)
}

func row(child *etree.Element) {
	fmt.Println("Processa Linha" + child.Tag)
}

func col(child *etree.Element) {
	fmt.Println("Processa Coluna" + child.Tag)
}

func label(child *etree.Element) {
	fmt.Println("Processa Label" + child.Tag)
}

func br(child *etree.Element) {
	fmt.Println("Processa Quebra de linha " + child.Tag)
}

func _barcode(child *etree.Element) {
	fmt.Println("Processa Codigo de barras " + child.Tag)
}
