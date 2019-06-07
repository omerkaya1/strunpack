package strunpack

import (
	"strconv"
	"strings"
	"unicode"
)

func StringParse(strToParse string) string {
	if strToParse == "" {
		return strToParse
	}
	str := strings.Builder{}
	toRune := []rune(strToParse)
	for i := 0; i < len(toRune); i++ {
		switch {
		// Letters
		case unicode.IsLetter(toRune[i]):
			str.WriteRune(toRune[i])
			break
		// Slashes
		case toRune[i] == []rune(`\`)[0]:
			str.WriteRune(toRune[i+1])
			i += 1
			break
		// Digits mixed with letters
		case unicode.IsDigit(toRune[i]) && i > 0 && str.Len() != 0:
			digits := make([]byte, 0)
			end := true
			for j := i; j < len(toRune); j++ {
				if !unicode.IsDigit(toRune[j]) {
					m, _ := strconv.ParseInt(string(digits), 10, 64)
					str.WriteString(strings.Repeat(string(str.String()[str.Len()-1]), int(m)-1))
					end = false
					i = j - 1
					break
				}
				digits = append(digits, byte(toRune[j]))
			}
			if end {
				m, _ := strconv.ParseInt(string(digits), 10, 64)
				str.WriteString(strings.Repeat(string(str.String()[str.Len()-1]), int(m)-1))
			}
			break
		// Ignore only digits
		default:
			break
		}
	}
	return str.String()
}
