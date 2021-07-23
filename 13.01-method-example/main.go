package main

import "fmt"

/*
	- Ngay: 22.07.2021
	- By: Nguyen Dang Duc
*/
func main() {
	// 01. Khai bao bien kieu Struct
	var (
		myPoint = Point{
			pointX0: 1.0,
			pointY0: 2.0,
		}
	)

	// In man hinh
	fmt.Println("Thong tin toa do: ", myPoint)

	isAbove := myPoint.IsAbove(0.0)
	fmt.Println("Tung do co nam tren duong y = 0.0 ? : ", isAbove)
}
