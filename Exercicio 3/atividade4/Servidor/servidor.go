package Servidor

import (
	"strings"
)

func RemoverVogais(w string) string {
	sv := w
	for _, c := range []string{"a", "e", "i", "o", "u"} {
		sv = strings.ReplaceAll(sv, c, "")
	}
	return sv
}
