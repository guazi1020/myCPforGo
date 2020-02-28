/*
本来想还需要转换为model,json才能用.没想到数组也可以。至少现在没有了。
*/

package ModelOperation

import (
	"fmt"
	"myCPforGo/Business/WebCralwer"
)

func ListToGameModel() {
	results := WebCralwer.SearchForGame("AC米兰", 10, 0)
	fmt.Println(results)
}
