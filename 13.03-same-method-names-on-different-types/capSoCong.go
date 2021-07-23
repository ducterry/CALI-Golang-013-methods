package main

type CapSoCong struct {
	soHangDau, congSai float64
}

func (csCong CapSoCong) NthTerm(soHangThuN int) float64 {
	return csCong.soHangDau + float64(soHangThuN-1)*csCong.congSai
}
