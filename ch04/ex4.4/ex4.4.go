// 4.4 Write a version of rotate that operates in a single pass.
package rotate

import "fmt"

func rotateLeft(a []int, numRot int) []int {
	lenA := len(a)
	if numRot <= 0 || numRot >= lenA {
		return a
	}
	temp := make([]int, lenA)
	for i, j := 0, numRot; i < lenA; i, j = i+1, j+1 {
		if j== len(a) {
			j = 0
		}
		temp[i] = a[j]
	}
	return temp
}

func main() {
	s := []int{1, 2, 4, 5}
	fmt.Println(rotateLeft(s,1))
}
