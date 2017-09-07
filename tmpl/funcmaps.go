package tmpl

import (
	"bytes"
	"html/template"
	"strings"
	"time"

	"strconv"

	"fmt"

	"html"

	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/util"
)

var funcMap = template.FuncMap{
	"today":                  today,
	"brdate":                 brDate,
	"replace":                replace,
	"docType":                docType,
	"trim":                   trim,
	"padLeft":                padLeft,
	"toString":               toString,
	"fmtDigitableLine":       fmtDigitableLine,
	"fmtCNPJ":                fmtCNPJ,
	"fmtCPF":                 fmtCPF,
	"fmtDoc":                 fmtDoc,
	"fmtNumber":              fmtNumber,
	"brDateWithoutDelimiter": brDateWithoutDelimiter,
	"enDateWithoutDelimiter": enDateWithoutDelimiter,
	"fullDate":               fulldate,
	"enDate":                 enDate,
	"hasErrorTags":           hasErrorTags,
	"toFloatStr":             toFloatStr,
	"concat":                 concat,
	"base64":                 base64,
	"unscape":                unscape,
	"unescapeHtmlString":     unescapeHtmlString,
}

func GetFuncMaps() template.FuncMap {
	return funcMap
}
func padLeft(value, char string, total uint) string {
	s := util.PadLeft(value, char, total)
	return s
}
func unscape(s string) template.HTML {
	return template.HTML(s)
}

func unescapeHtmlString(s string) template.HTML {
	return template.HTML(html.UnescapeString(s))
}

func hasErrorTags(mapValues map[string]string, errorTags ...string) bool {
	hasError := false
	for _, v := range errorTags {
		if value, exist := mapValues[v]; exist && strings.Trim(value, " ") != "" {
			hasError = true
			break
		}
	}
	return hasError
}

func fmtNumber(n uint64) string {
	real := n / 100
	cents := n % 100
	return fmt.Sprintf("%d,%02d", real, cents)
}

func toFloatStr(n uint64) string {
	real := n / 100
	cents := n % 100
	return fmt.Sprintf("%d.%02d", real, cents)
}

func fmtDoc(doc models.Document) string {
	if e := doc.ValidateCPF(); e == nil {
		return fmtCPF(doc.Number)
	}
	return fmtCNPJ(doc.Number)
}

func toString(number uint) string {
	return strconv.FormatInt(int64(number), 10)
}

func today() time.Time {
	return time.Now()
}

func fulldate(t time.Time) string {
	return t.Format("20060102150405")
}

func brDate(d time.Time) string {
	return d.Format("02/01/2006")
}

func enDate(d time.Time, del string) string {
	return d.Format("2006" + del + "01" + del + "02")
}

func brDateWithoutDelimiter(d time.Time) string {
	return d.Format("02012006")
}

func enDateWithoutDelimiter(d time.Time) string {
	return d.Format("20060102")
}

func replace(str, old, new string) string {
	return strings.Replace(str, old, new, -1)
}

func docType(s models.Document) int {
	if s.IsCPF() {
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

func concat(s ...string) string {
	buf := bytes.Buffer{}
	for _, item := range s {
		buf.WriteString(item)
	}
	return buf.String()
}

func base64(s string) string {
	return util.Base64(s)
}
