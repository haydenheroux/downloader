package main

import (
	"strings"
)

func ChangeExtension(fileName string, extension string) string {
	temp := strings.Split(fileName, ".")
	str := strings.Join(temp[:len(temp)-1], ".")
	return str + "." + extension
}
