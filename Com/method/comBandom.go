package baseMethod

/*
* author:guazi1020
* create_date:2019/10
* describe:随机数生成
 */
import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//Bandom 随机数实体
//Number
//Min 最小数
//Max 最大数
type Bandom struct {
	Number int
	Min    int
	Max    int
}

//CreatRandomInt() 生成Int随机数
func (bandom *Bandom) CreatRandomInt() int {
	seed := time.Now().UnixNano()
	//rand 随机数
	r := rand.New(rand.NewSource(seed))
	return (r.Intn(bandom.Max-bandom.Min) + bandom.Min)

}

//生成float64随机数(m:数据保留数量)
func (bandom *Bandom) CreatRandomFloat(m int) float64 {
	seed := time.Now().UnixNano()
	//rand 随机数
	r := rand.New(rand.NewSource(seed))
	// float_num, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(r.Intn(bandom.Max-bandom.Min))+r.Float64()), 64)
	float_num, _ := strconv.ParseFloat(ChangeNumber(float64(r.Intn(bandom.Max-bandom.Min))+r.Float64(), m), 64)
	return float_num
}

/*
保留小数点后几位方法(f:目标数据,m:数据保留位数)
*/
func ChangeNumber(f float64, m int) string {
	n := strconv.FormatFloat(f, 'f', -1, 32)
	if n == "" {
		return ""
	}
	if m >= len(n) {
		return n
	}
	newn := strings.Split(n, ".")
	if len(newn) < 2 || m >= len(newn[1]) {
		return n
	}
	return newn[0] + "." + newn[1][:m]
}
