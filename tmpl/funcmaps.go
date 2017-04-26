package tmpl

import (
	"bytes"
	"html/template"
	"strings"
	"time"

	"strconv"

	"bitbucket.org/mundipagg/boletoapi/models"
	"bitbucket.org/mundipagg/boletoapi/util"
)

var funcMap = template.FuncMap{
	"today":            today,
	"brdate":           brDate,
	"replace":          replace,
	"docType":          docType,
	"trim":             trim,
	"padLeft":          util.PadLeft,
	"toString":         toString,
	"fmtDigitableLine": fmtDigitableLine,
	"fmtCNPJ":          fmtCNPJ,
	"fmtCPF":           fmtCPF,
	"attr": func(s string) template.HTMLAttr {
		return template.HTMLAttr(s)
	},
}

func toString(number int) string {
	return strconv.Itoa(number)
}

func today() time.Time {
	return time.Now()
}

func brDate(d time.Time) string {
	return d.Format("02/01/2006")
}

func replace(str, old, new string) string {
	return strings.Replace(str, old, new, -1)
}

func docType(s models.DocumentType) int {
	if s.IsCpf() {
		return 1
	}
	return 2
}

func trim(s string) string {
	return strings.TrimSpace(s)
}
func fmtDigitableLine(s string) string {
	buf := bytes.Buffer{}
	for idx, c := range s {
		if idx == 5 || idx == 15 || idx == 26 {
			buf.WriteString(".")
		}
		if idx == 10 || idx == 21 || idx == 32 || idx == 33 {
			buf.WriteString(" ")
		}
		buf.WriteByte(byte(c))
	}
	return buf.String()
}

func fmtCNPJ(s string) string {
	buf := bytes.Buffer{}
	for idx, c := range s {
		if idx == 2 || idx == 5 {
			buf.WriteString(".")
		}
		if idx == 8 {
			buf.WriteString("/")
		}
		if idx == 12 {
			buf.WriteString("-")
		}
		buf.WriteRune(c)
	}
	return buf.String()
}

func fmtCPF(s string) string {
	buf := bytes.Buffer{}
	for idx, c := range s {
		if idx == 3 || idx == 6 {
			buf.WriteString(".")
		}
		if idx == 9 {
			buf.WriteString("-")
		}
		buf.WriteRune(c)
	}
	return buf.String()
}
