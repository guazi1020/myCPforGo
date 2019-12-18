package Test

import (
	"myCP/Com"
	"testing"
)

func TestHello(t *testing.T) {
	if "sdfas" != Com.RemoveBlank("sdf  as  ") {
		t.Error("erro")
	} else {
		t.Log("it's ok")
	}

	//t.Log()
	//t.Log("Hello World")
}
