package tmpl

import (
	"html/template"
	"strings"
	"time"

	"fmt"

	"strconv"

	"bitbucket.org/mundipagg/boletoapi/models"
)

var funcMap = template.FuncMap{
	"today":   today,
	"brdate":  brDate,
	"replace": replace,
	"docType": docType,
	"trim":    trim,
	"padLeft": padLeft,
}

func padLeft(value, char string, total uint) string {
	s := "%" + char + strconv.Itoa(int(total)) + "s"
	return fmt.Sprintf(s, value)
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
