package main

import "math"

type CapSoNhan struct {
	soHangDau, congBoi float64
}

func (csNhan CapSoNhan) NthTerm(soHangThuN int) float64 {
	return csNhan.soHangDau * math.Pow(csNhan.congBoi, float64(soHangThuN-1))
}
