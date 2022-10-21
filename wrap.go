package gg

import (
	"fmt"
	"strings"
)

type measureStringer interface {
	MeasureString(s string) (w, h float64)
}

func splitRune(x string) []string {
	var result []string
	for _, c := range strings.Split(x, "") {
		result = append(result, string(c))
	}
	fmt.Println("2--", result)
	return result
}

func wordWrap(m measureStringer, s string, width float64) []string {
	var result []string
	for _, line := range strings.Split(s, "\n") {
		fmt.Println("1--", line)
		fields := splitRune(line)
		x := ""
		for i := 0; i < len(fields); i++ {
			w, _ := m.MeasureString(x + fields[i])
			if w > width {
				if x == "" {
					result = append(result, fields[i])
					x = ""
					continue
				} else {
					result = append(result, x)
					x = ""
				}
			}
			x += fields[i]
		}
		if x != "" {
			result = append(result, x)
		}
	}
	fmt.Println("3--", result)
	return result
}
