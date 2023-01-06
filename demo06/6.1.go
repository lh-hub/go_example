// 简单的 return 语句也可以用来结束 for 死循环，或者结束一个协程（goroutine）
package main

func main() {
	println("In main before calling greeting")
	greeting()
	println("In main after calling greeting")
}

func greeting() {
	println("In greeting: Hi!!!!!")
}
