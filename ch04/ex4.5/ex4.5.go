package main

import "fmt"

func removeDup(strings []string) []string {
	length := len(strings)
	skipped := 0
	for i := 0; i < length; {
		// Find ajacent duplicates
		findDup := false
		j := i + 1
		for ; j < length; j++ {
			if strings[i] != strings[j] {
				break
			} else {
				findDup = true
				skipped++
			}
		}

		if findDup {
			copy(strings[i:], strings[j-1:])
		}
		i = j
	}
	return strings[:length-skipped]
}

func main() {
	data := []string{"one", "one", "one", "three", "three"}
	fmt.Printf("%q\n", removeDup(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)
}
