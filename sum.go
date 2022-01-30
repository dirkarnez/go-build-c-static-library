// package sum

// func (a, b int) int {
// 	return a + b
// }

package main

import "fmt"
import "C"

func main() {}

//export Sum
func Sum(a, b int) int {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	return a + b
}
