package util

import (
	"fmt"
	"strconv"
	"unicode"

	"encoding/json"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

//RemoveDiacritics remove caracteres especiais de uma string
func RemoveDiacritics(s string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

// PadLeft insere um caractere a esquerda de um texto
func PadLeft(value, char string, total uint) string {
	s := "%" + char + strconv.Itoa(int(total)) + "s"
	return fmt.Sprintf(s, value)
}

//Stringify convete objeto para JSON
func Stringify(o interface{}) string {
	b, _ := json.Marshal(o)
	return string(b)
}

//ParseJSON converte string para um objeto GO
func ParseJSON(s string, o interface{}) interface{} {
	json.Unmarshal([]byte(s), o)
	return o
}
