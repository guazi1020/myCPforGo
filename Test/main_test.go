package Test

import (
	"myCPforGo/Business/CPHttp/ImpMethod"
	"myCPforGo/Business/WebCralwer"
	"myCPforGo/Com/comerr"
	"myCPforGo/Interface/HTTP"
	"testing"
)

func TestHello(t *testing.T) {

	WebCralwer.FindEInfoByUUID("097dd210-d8a2-a217-e17c-6d503925d40d")
	// str := ""
	// stri, _ := strconv.Atoi(str)
	// str1 := ""
	// str1i, _ := strconv.Atoi(str1)
	// if stri < str1i {
	// 	t.Log("<")
	// } else {
	// 	t.Log(">")
	// }
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
func TestInterface(t *testing.T) {
	comerr.ReadConfig()
}

func Testtt(t *testing.T) {
	var vv HTTP.IGetGameData
	vv = ImpMethod.GetGameDataOne{}
	vv.GetGameDataForYear("2019")
}
