package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var count int

	//	for i := 0; i < 64; i++ {
	//		count += int(pc[byte(x>>(uint64(i)*8))])
	//	}

	for i := 0; i < 64; i++ {
		count += int(x>>uint(i)) & 1
	}
	return count
}
