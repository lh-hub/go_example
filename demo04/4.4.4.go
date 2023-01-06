package main

func init() {
	Pi := 3.14
	println(Pi)
}

func main() {
	a, b, c := 5, 7, "abc"
	a, b = b, a
	print(a, b, c)
}

// 但是全局变量是允许声明但不使用。
