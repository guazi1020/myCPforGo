package Method

import (
	"fmt"
	"html"
	"myCPforGo/Business/CPHttp/ImpMethod"
	"myCPforGo/Interface/HTTP"
	"net/http"
)

//HandleIndex /index
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Router test: hello,%q", html.EscapeString(r.URL.Path))
	var getGame HTTP.IGetGameData
	getGame = ImpMethod.GetGameDataOne{Year: "2020"}
	getGame.GetGameDataForYear()
}
