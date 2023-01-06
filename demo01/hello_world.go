package main

func main() {
	println("hello", "world")
}

// gofmt -r "(hell) -> helloo000" -w .\hello_world.go 无效
// gofmt -r "mainn -> main" -w .\hello_world.go 有效
