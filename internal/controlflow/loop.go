package controlflow

import "fmt"

func RunLoop() {
	// for ประกาศตัวแปร; เงื่อนไข; การเพิ่ม(ลด)ค่า
	for i := 0; i <= 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	c := 3
	// for เงื่อนไข
	for c > 0 {
		fmt.Print(c, " ")
		// c = c - 1
		c--
	}
	fmt.Println()

	// for with range
	d := []int{0, 1, 1, 2, 3, 5, 8}
	for _, value := range d {
		fmt.Print(value, " ")
	}
	fmt.Println()
}
