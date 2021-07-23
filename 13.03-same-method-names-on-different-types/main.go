package main

import "fmt"

/*
	- Ngay: 23.07.2021
	- By: Nguyen Dang Duc
*/

func main() {
	// 01. Khai báo biến cấu trúc
	var (
		capSoCong = CapSoCong{
			congSai:   10,
			soHangDau: 1}

		capSoNhan = CapSoNhan{
			congBoi:   5,
			soHangDau: 1}
	)

	// 02. Tính toán
	soHangThu5CSC := capSoCong.NthTerm(5)
	soHangThu5CSN := capSoNhan.NthTerm(5)

	fmt.Println("Số hạng thứ 5 của cấp số cộng là : ", soHangThu5CSC)
	fmt.Println("Số hạng thứ 5 của cấp số nhân là : ", soHangThu5CSN)
}
