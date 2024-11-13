package unpacker

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(str string) (string, error) {
	rs := []rune(str)
	var result strings.Builder

	var count int
	var isEscape bool

	for i, v := range rs {
		if isEscape {
			result.WriteRune(v)
			isEscape = false
		} else if unicode.IsDigit(v) {
			count++
			num, err := strconv.Atoi(string(v))
			if err != nil {
				return "", err
			}

			if i > 0 {
				result.WriteString(strings.Repeat(string(rs[i-1]), num-1))
			} else {
				result.WriteRune(v)
			}
		} else if string(v) == `\` {
			isEscape = true
		} else {
			result.WriteRune(v)
		}
	}

	if str == "" {
		return "", nil
	} else if len(rs) == count {
		return "", fmt.Errorf("Некорректная строка")
	} else {
		return result.String(), nil
	}
}
