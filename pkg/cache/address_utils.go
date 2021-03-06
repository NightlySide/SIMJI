package cache

// DecodeAddress prend une adresse et retourne sa version
// décodée
func DecodeAddress(addr int) (int, int, int) {
	var tag int
	var setID int
	var blockOffset int

	tag			= (addr & 0xFFFFF000) >> (4*3)
	setID	    = (addr & 0x00000FF0) >> (4*1)
	blockOffset = (addr & 0x0000000F)

	return tag, setID, blockOffset
}

// EncodeAddress retourne une adresse à partir de ses éléments
// (tag, setID, blockoffset)
func EncodeAddress(tag int, setID int, blockOffset int) int {
	var res int
	res += tag << (4*3)
	res += setID << (4*1)
	res += blockOffset

	return res
}

// AddressFromIndex retourne une addresse en fonction de l'indice de la position
// de la variable, ou bien de l'indice du registre
func (c *Cache) AddressFromIndex(index int) int {
	var tag, setID, blockOffset int

	setID = (index / c.wordsNb ) % c.setsNb
	tag = index / c.wordsNb
	blockOffset = index % c.wordsNb

	return EncodeAddress(tag, setID, blockOffset)
}
