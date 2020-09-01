package main

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {

	str := `abc:ab"abc"`

	t.Log(str)

	str1 := `abc{cc}adf`
	str11 := strings.Replace(str1, "{cc}", "111", 1)
	t.Log(str11)

	str2 := `abc"abc"abc`
	t.Log(strings.ReplaceAll(str2, "\"", ""))
}
