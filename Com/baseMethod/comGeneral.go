package baseMethod

//Compoundrate 计算复利
//
func Compoundrate(source float64, rate float64, number int) float64 {

	for i := 0; i < number; i++ {
		source = source * (1 + rate)
	}

	return source
}

//CountMultiplying 幂次方
func CountMultiplying(source float64, number int) float64 {
	count := number
	if number < 0 {
		count = -number
	}
	for i := 0; i < count; i++ {
		source = source * source
	}
	if number < 0 {
		source = 1 / source
	}
	return source
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
