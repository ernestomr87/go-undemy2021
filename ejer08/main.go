package main

import "fmt"

var table [10]int
var matrix [5][7]int

func main() {
	// table[0] = 1
	// table[5] = 45
	// fmt.Println(table)

	// table1 := [10]int{1, 2, 3, 4}
	// for _, item := range table1 {
	// 	fmt.Println(item)

	// }

	// matrix[3][5] = 1
	// fmt.Println(matrix)

	matrix := []int{2, 3, 4, 5}
	fmt.Println(matrix)

	matrix = append(matrix, 9)
	fmt.Println(matrix)
	variant2()
	variant3()
	variant4()
}

func variant2() {
	elements := [5]int{1, 2, 3, 4, 5}
	// portion := elements[3:]
	// portion := elements[:4]
	portion := elements[2:4]
	fmt.Println(portion)
}

func variant3() {
	elements := make([]int, 5, 20)
	fmt.Printf("Largest %d, Capacity %d\n", len(elements), cap(elements))
}

func variant4() {
	nums := make([]int, 0, 0)
	fmt.Printf("Largest %d, Capacity %d\n", len(nums), cap(nums))

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largest %d, Capacity %d\n", len(nums), cap(nums))
}
