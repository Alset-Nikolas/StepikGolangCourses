package main

import "fmt"

func main() {
	var d int
	var h int
	var m int

	fmt.Scan(&d)

	h = d / (360 / 12)
	m = 2 * (d - h*(360/12))

	fmt.Print(fmt.Sprintf("It is %d hours %d minutes.", h, m))

}
