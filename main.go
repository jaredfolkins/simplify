package main

import "fmt"

var a = []int{0, 9, 8, 3, 3, 3, 5, 9, 5, 0}
var s = make([]int, len(a)/2+1)

func main() {

	fmt.Println("before:", a)

	mergeSort(a)
	a := removeDupes(a)

	fmt.Println("after: ", a)

}

func mergeSort(a []int) {
	if len(a) < 2 {
		return
	}

	mid := len(a) / 2
	mergeSort(a[:mid])
	mergeSort(a[mid:])
	if a[mid-1] <= a[mid] {
		return

	}

	copy(s, a[:mid])
	l, r := 0, mid
	for i := 0; ; i++ {
		if s[l] <= a[r] {
			a[i] = s[l]
			l++
			if l == mid {
				break

			}

		} else {
			a[i] = a[r]
			r++
			if r == len(a) {
				copy(a[i+1:], s[l:mid])
				break

			}

		}

	}
	return

}

func removeDupes(a []int) []int {
	result := []int{}
	seen := map[int]int{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}
