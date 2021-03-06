package baseMethod

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

//阶乘计算，x底数 n阶乘数
func MyPow(x float64, n int) float64 {
	if n < 0 {
		if n == -1<<31 {
			return MyPow(x, n+1) / x
		}
		x = 1 / x
		n *= -1
	}

	var result float64 = 1
	var current_product float64 = x
	for n > 0 {
		if n&1 == 1 { // 当前位为1的话就乘上去
			result *= current_product
		}
		// 前进一位
		current_product *= current_product
		n = n >> 1
	}
	return result
}

//小数变分数
func DecimalsToGrade(source float64) (int64, int64) {

	var denominator int64 //denominator 分母
	var numerator int64   //分子
	s1 := strconv.FormatFloat(source, 'f', -1, 64)
	strs := strings.Split(s1, ".")
	if len(strs) > 1 {
		//将分子去小数化
		_num := source
		_den := 1
		for i := 0; i < len(strs[1]); i++ {
			_den = _den * 10
			_num = _num * 10
		}

		var i int64
		//list_x := make([]int, int(_den))
		list_x := []int64{}
		list_y := []int64{}

		for i = 1; i <= int64(_den); i++ {
			v := decimal.NewFromInt(int64(_den)).Mod(decimal.NewFromInt(i))
			if v.IntPart() == 0 {
				list_x = append(list_x, i)
			}
		}
		for i = 1; i <= int64(_num); i++ {
			v := decimal.NewFromInt(int64(_num)).Mod(decimal.NewFromInt(i))
			if v.IntPart() == 0 {
				list_y = append(list_y, i)
			}
		}
		var flag int64
		flag = 1
		for i := 0; i < len(list_x); i++ {
			for j := 0; j < len(list_y); j++ {
				if list_x[i] == list_y[j] {
					flag = list_x[i]
				}
			}
		}
		//	denominator =
		denominator = decimal.NewFromInt(int64(_den)).Div(decimal.NewFromInt(flag)).IntPart()
		numerator = decimal.NewFromInt(int64(_num)).Div(decimal.NewFromInt(flag)).IntPart()
		//fmt.Println(numerator, "/", denominator)
		//fmt.Println(list_x, list_y)

	} else {
		denominator = 1
		numerator = int64(source)
	}
	return numerator, denominator
}

//Compoundrate 计算复利
//source 本 rate利率
func Compoundrate(source float64, rate float64, number int) float64 {

	for i := 0; i < number; i++ {
		source = source * (1 + rate)
	}

	return source
}

//CountMutiplyingsqrt 幂次方，带小数的
//source 数据源	number 幂次
func CountMultiplyingsqrt(source float64, number float64) float64 {
	s1 := strconv.FormatFloat(number, 'f', -1, 64)
	strs := strings.Split(s1, ".")
	if len(strs) > 1 { //如果有小数
		//CountMultiplying(strconv.Atoi(strs[1]))
		fmt.Println(len(strs[1]))
	} else { //如果没有小数直接用整数方法
		_number, _ := strconv.Atoi(strs[0])
		fmt.Println(CountMultiplying(source, _number))
	}

	return source
}

//CountMultiplying 幂次方、
//source 源数据 number 幂次整数
func CountMultiplying(source float64, number int) float64 {
	count := number
	var result float64
	result = 1
	if number < 0 {
		count = -number
	}
	for i := 0; i < count; i++ {
		result = source * result
		//fmt.Println("source", result)
	}
	if number < 0 {
		result = 1 / result
	}
	return result
}

//CountFactorial 阶乘 3!
func CountFactorial(source int) int {
	var result int
	result = 1
	if source > 0 {
		result = source * CountFactorial(source-1)
	}
	return result
}
