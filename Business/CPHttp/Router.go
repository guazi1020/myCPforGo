package CPHttp

import (
	"fmt"
	"html"
	"log"
	"myCPforGo/Business/CPHttp/ImpMethod"
	"myCPforGo/Interface/HTTP"
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
	router.HandleFunc("/", HandleIndex)

	err := http.ListenAndServe(_port, router) //监听端口,装载路由
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//HandleIndex /index
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Router test: hello,%q", html.EscapeString(r.URL.Path))
	var getGame HTTP.IGetGameData
	getGame = ImpMethod.GetGameDataOne{}
	getGame.GetGameDataForYear()
}
