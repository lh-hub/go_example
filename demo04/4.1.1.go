package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	println(a)
}

func m() {
	// a := "O"
	a = "O"
	println(a)
}
