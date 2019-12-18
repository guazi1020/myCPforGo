package comerr

//CheckErr 容错
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
