package Test

import (
	"strconv"
	"testing"
)

func TestHello(t *testing.T) {
	str := ""
	stri, _ := strconv.Atoi(str)
	str1 := ""
	str1i, _ := strconv.Atoi(str1)
	if stri < str1i {
		t.Log("<")
	} else {
		t.Log(">")
	}
	// str := "[123]"
	// if len(str) != 0 {
	// 	str = str[1 : len(str)-1]
	// }

	// fmt.Println(str)

	// if "sdfas" != Com.RemoveBlank("sdf  as  ") {
	// 	t.Error("erro")
	// } else {
	// 	t.Log("it's ok")
	// }

}
