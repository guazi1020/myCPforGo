package CPHttp

import (
	"log"
	"myCPforGo/Business/CPHttp/Method"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	_port = ":8080"
)

//StartHttp 开始启动httpweb
func StartHttp() {
	router := mux.NewRouter().StrictSlash(true)
	//注册
	router.HandleFunc("/", Method.HandleIndex)

	err := http.ListenAndServe(_port, router) //监听端口,装载路由
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
