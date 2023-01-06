package main

func main() {
	var aVar = 10
	println(aVar == 5)
	println(aVar == 10)

	println(aVar != 5)
	println(aVar != 10)
}
// Go 对于值之间的比较有非常严格的限制，只有两个类型相同的值才可以进行比较