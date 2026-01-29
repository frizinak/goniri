package generate

import (
	"strings"
	"unicode"
)

func Camel2Dash(i string) string {
	return camel2(i, '-')
}

func Dash2Camel(i string) string {
	p := strings.FieldsFunc(i, func(r rune) bool {
		return r == '-' || r == '_'
	})
	for i := 1; i < len(p); i++ {
		p[i] = UCFirst(strings.ToLower(p[i]))
	}
	return strings.Join(p, "")
}

func UCFirst(i string) string {
	n := []rune(i)
	n[0] = unicode.ToUpper(n[0])
	return string(n)
}

func camel2(i string, char rune) string {
	a := []rune(i)
	b := make([]rune, 0, len(a)+5)
	for i, ch := range a {
		if unicode.IsUpper(ch) {
			if i != 0 {
				b = append(b, char)
			}
			ch = unicode.ToLower(ch)
		}
		b = append(b, ch)
	}
	return string(b)
}
