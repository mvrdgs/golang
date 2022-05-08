package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	texto := "Maurício tem 31 anos"
	expr := regexp.MustCompile("\\d")

	fmt.Println(expr.ReplaceAllString(texto, "2"))

	expr2 := regexp.MustCompile("\\b\\w")
	texto2 := "antonio carlos jobim"

	transformCase := func(s string) string { return strings.ToUpper(s) }

	fmt.Println(transformCase(texto2))
	processado := expr2.ReplaceAllStringFunc(texto2, transformCase)

	fmt.Println(processado)
}
