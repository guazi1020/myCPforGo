package CPHttp

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func StartHttp() {
// 	router := mux.NewRouter().StrictSlash(true)
// 	//注册
// 	// router.HandleFunc("/", HandleIndex)
// 	// router.HandleFunc("/app", HandleDemoIndex)
// 	// router.HandleFunc("/app/{id}", HandleDemoShow)
// 	// router.HandleFunc("/SaveResultDate/{begindate}/{enddate}", SaveResultDate)
// 	// router.HandleFunc("/GetNowGame", GetNowGame)

// 	err := http.ListenAndServe(_port, router) //监听端口,装载路由
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
