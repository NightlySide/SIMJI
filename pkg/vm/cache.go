package vm

//-----------------------------//
// (S, E, B, m)
// S : 2^s number of sets
// E : number of lines per set (here E=1)
// B : 2^b bloc_size (byte)
// m = log2(M) number of physical (main memory) address bits
//----------------------------//
type Cache struct {
	sets []Set
	memory *Memory
}

func NewCache(nbSets int, nbLines int, nbMots int, sizeMot int) *Cache {
	c := &Cache{}
	for i := 0; i < nbSets; i++ {
		c.sets = append(c.sets, *NewSet(nbLines, nbMots, sizeMot))
	}
	c.memory = NewMemory(1024)
	return c
}

func (c *Cache) InCache(addr int) (bool, Line) {
	tag, setID, _ := decodeAddress(addr)

	set := c.sets[setID]
	for _, line := range set.lines {
		if line.valid == 1 {
			if line.tag == tag {
				return true, line
			}
		}
	}

	return false, Line{}
}

func (c *Cache) Read(addr int) int {
	isInCache, line := c.InCache(addr)
	_, _, blockOffset := decodeAddress(addr)

	if isInCache {
		return line.data[blockOffset].data
	} else {
		mot := c.memory.Read(addr)
		c.Write(addr, mot)
		return mot
	}
}

func (c *Cache) Write(addr int, data int) {

}

func (c *Cache) WriteThrough(addr int, data int) {
	c.Write(addr, data)
	c.memory.Write(addr, data)
}

func decodeAddress(addr int) (int, int, int) {
	var tag int
	var setId int
	var blockOffset int

	tag			= (addr & 0xFFFFF000) >> (4*3)
	setId	    = (addr & 0x00000FF0) >> (4*1) 
	blockOffset = (addr & 0x0000000F)

	return tag, setId, blockOffset
}

// --------------- SET --------------

type Set struct {
	lines []Line
}

func NewSet(nbLines int, nbMots int, sizeMot int) *Set {
	s := &Set{}
	for i := 0; i < nbLines; i++ {
		s.lines = append(s.lines, *NewLine(nbMots, sizeMot))
	}
	return s
}

// --------------- LINE --------------

type Line struct {
	valid 	int
	tag 	int
	data 	[]Mot
	nbMots 	int 
}

func NewLine(nbMots int, sizeMot int) *Line {
	l := &Line{}
	for i := 0; i < nbMots; i++ {
		l.data = append(l.data, Mot{size: sizeMot})
	}
	return l
}

// --------------- MOT --------------

type Mot struct {
	size int
	data int
}