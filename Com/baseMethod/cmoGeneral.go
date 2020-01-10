package baseMethod

//Compoundrate 计算复利
//
func Compoundrate(source float64, rate float64, number int) float64 {

	for i := 0; i < number; i++ {
		source = source * (1 + rate)
	}

	return source
}
