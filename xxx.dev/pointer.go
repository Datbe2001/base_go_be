package main

import "fmt"

func main() {
	a := 10 // Biến bình thường
	p := &a // Con trỏ trỏ tới địa chỉ của a

	fmt.Println("Giá trị của a:", a)  // 10
	fmt.Println("Địa chỉ của a:", &a) // Ví dụ: 0xc000018090
	fmt.Println("Con trỏ p:", p)      // 0xc000018090
	fmt.Println("Giá trị qua p:", *p) // 10

	*p = 20                               // Thay đổi giá trị qua con trỏ
	fmt.Println("Giá trị mới của a:", a)  // 20
	fmt.Println("Giá trị mới của p:", *p) // 20
}
