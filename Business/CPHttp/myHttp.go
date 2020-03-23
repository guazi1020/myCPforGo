package CPHttp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"myCPforGo/Business/WebCralwer"
	"myCPforGo/Model"
	"net/http"
	"os"

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
	router.HandleFunc("/app", HandleDemoIndex)
	router.HandleFunc("/app/{id}", HandleDemoShow)
	router.HandleFunc("/dataforjson", DataForJson)

	err := http.ListenAndServe(_port, router) //监听端口,装载路由
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func DataForJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")         //header的类型
	w.Header().Set("content-type", "application/json")                     //返回数据格式是json
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	str, _ := os.Getwd()
	inputFile, inputError := os.Open(str + "\\Document\\result_data1.json")
	if inputError != nil {
		fmt.Println(inputError)
		return
	}
	defer inputFile.Close()

	var s string
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break
		}
		s = s + inputString
	}
	fmt.Fprintf(w, s)
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
