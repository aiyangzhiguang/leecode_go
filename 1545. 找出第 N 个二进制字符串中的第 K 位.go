package main

func findKthBit(n int, k int) byte {
	//var res []byte
	if k == 1 {
		return '0'
	}
	mid := 1 << (n - 1)
	if k == mid {
		return '1'
	} else if k < mid {
		return findKthBit(n-1, k)
	} else {
		return invertByte(findKthBit(n-1, mid*2-k))
	}
}

func invertByte(b byte) byte {
	return '0' + '1' - b
}
