package cache

func DecodeAddress(addr int) (int, int, int) {
	var tag int
	var setID int
	var blockOffset int

	tag			= (addr & 0xFFFFF000) >> (4*3)
	setID	    = (addr & 0x00000FF0) >> (4*1)
	blockOffset = (addr & 0x0000000F)

	return tag, setID, blockOffset
}

func EncodeAddress(tag int, setID int, blockOffset int) int {
	var res int
	res += tag << (4*3)
	res += setID << (4*1)
	res += blockOffset

	return res
}