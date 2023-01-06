package main

func main() {
	var i int = 5
	for {
		i = i - 1
		println(i)
		if i < 0 {
			break
		}
	}
}
