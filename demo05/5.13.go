package main

func main() {

LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			println(j)
		}
	}
}
