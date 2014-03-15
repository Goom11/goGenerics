package main

import (
	"fmt"
	"strings"
)

var types = []string{
	"int",
	"float64",
	"string",
}

func genFoldLeft(t1 string, t2 string) {
	foldLeftTemplate := fmt.Sprintf("func FoldLeft%s%s(l []%s, z %s, f func(%s, %s) %s) %s {\n\tacc := z\n\tthese := l\n\tfor len(these) != 0 {\n\t\tacc = f(acc, these[0])\n\t\tthese = these[1:]\n\t}\n\treturn acc\n}\n", t1, t2, t1, t2, t2, t1, t2, t2)
	fmt.Println(foldLeftTemplate)
}

func genSelfGen(t string) {
	var selfTemplate = []string{
		"func Self",
		"(x ",
		") ",
		" {\n\treturn x\n}\n",
	}
	fmt.Println(strings.Join(selfTemplate, t))
}

func main() {
	fmt.Println("package generics")
	fmt.Println("")
	for _, t := range types {
		genSelfGen(t)
	}
	for _, t1 := range types {
		for _, t2 := range types {
			if t1 != t2 {
				genFoldLeft(t1, t2)
			}
		}
	}
	for _, t := range types {
		genFoldLeft(t, t)
	}
}
