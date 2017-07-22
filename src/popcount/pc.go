package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountTable returns the population count (number of set bits) of x.
func PopCountTable(x uint64) int {
	return int(pc[byte(x>>0)] +
		pc[byte(x>>8)] +
		pc[byte(x>>16)] +
		pc[byte(x>>24)] +
		pc[byte(x>>32)] +
		pc[byte(x>>40)] +
		pc[byte(x>>48)] +
		pc[byte(x>>56)])
}

// PopCountLoop loops through each byte in x and adds up the count of each byte.
func PopCountLoop(x uint64) (count int) {
	for i := 0; i < 8; i++ {
		count += int(pc[(x>>uint(i*8))&0xFF])
	}
	return
}

// PopCountShift iterates over every bit of x and adds up the count of 1 bits.
func PopCountShift(x uint64) (count int) {
	for i := uint(0); i < 64; i++ {
		count += int((x >> i) & 1)
	}
	return
}

// PopCountClear iteratively clears the bottom bit of x, counting the numbe of times it can do this.
func PopCountClear(x uint64) (count int) {
	for x > 0 {
		count++
		x &= (x - 1)
	}
	return
}
