//Ex 4.4: Write a version of rotate that operates in place
package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4}
	//b := a[0:4]
	rotate(a, 2)
}

//rotates slice s left by n places
func leftRotate(s []int, n int) {

	r := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[(i+n)%len(s)]
	}
	fmt.Println(r)
}
