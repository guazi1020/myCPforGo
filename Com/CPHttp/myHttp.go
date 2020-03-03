package CPHttp

import (
	"encoding/json"
	"fmt"
	"html"
	"myCPforGo/Business/WebCralwer"
	"myCPforGo/Model"
	"net/http"

	"github.com/gorilla/mux"
)

//StartHttp 开始启动httpweb
func StartHttp() {
	router := mux.NewRouter().StrictSlash(true)
	//注册
	router.HandleFunc("/", HandleIndex)
	router.HandleFunc("/app", HandleDemoIndex)
	router.HandleFunc("/app/{id}", HandleDemoShow)

	fmt.Println("Main task")
	http.ListenAndServe(":8080", router)
	//log.Fatal(http.ListenAndServe(":8080", router))
}

//HandleIndex /index
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Router test: hello,%q", html.EscapeString(r.URL.Path))
}

//HandleDemoIndex /app
func HandleDemoIndex(w http.ResponseWriter, r *http.Request) {
	domains := Games{
		Model.Game{UUID: "a"},
		Model.Game{UUID: "b"},
	}
	_ = domains
	json.NewEncoder(w).Encode(WebCralwer.SearchForGame("AC米兰", 10, 0))
}

// fmt.Fprintf(w, "Router test: hello,%q", html.EscapeString(r.URL.Path))
// fmt.Fprintf(w, "this is app")
type Games []Model.Game

//HandDemoShow /app/{i}
func HandleDemoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Domain Show:%q", id)
}
