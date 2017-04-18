package boleto

import gin "gopkg.in/gin-gonic/gin.v1"
import "bitbucket.org/mundipagg/boletoapi/tmpl"

const template = `
	<html>
		<body>			
			<p>Ola mundo</p> 
		</body>
	</html>
`

func HTML(w gin.ResponseWriter, boleto interface{}) {
	b := tmpl.New()
	page, _ := b.From(boleto).To(template).Transform()
	w.WriteString(page)
}
