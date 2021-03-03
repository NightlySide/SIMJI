package cache

import "github.com/rs/zerolog/log"

// Cache is a structure containing memory cache data
type Cache struct {
	sets []Set
	memory *Memory
	setsNb int
	linesNb int
	wordSize int
	wordsNb int
}

// NewCache creates a new cache using the (S, E, B, m) notation
// S : 2^s number of sets
// E : number of lines per set (here E=1)
// B : 2^b bloc_size (byte)
// m = log2(M) number of physical (main memory) address bits
func NewCache(setsNb int, linesNb int, blocSize int, wordSize int) *Cache {
	// nbMots formula : m = t + s + b => 2^m = 2^(t+s+b) = 2^t*2^s+2^b
	// ie. nbMots = M = 2^t * S * B with t = 20

	c := &Cache{}
	c.wordsNb = blocSize*8 / wordSize
	c.setsNb = setsNb
	c.linesNb = linesNb
	c.wordSize = wordSize
	for i := 0; i < setsNb; i++ {
		c.sets = append(c.sets, *NewSet(c.linesNb, c.wordsNb, c.wordSize))
	}
	c.memory = NewMemory(1024, blocSize)
	return c
}

// InCache checks wether an address is already stored in the cache
func (c *Cache) InCache(addr int) (bool, Line) {
	tag, setID, _ := DecodeAddress(addr)

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

// Read read data from the cache
// If the data is not in the cache, it gets it from the memory
func (c *Cache) Read(addr int) int {
	isInCache, line := c.InCache(addr)
	tag, setID, blockOffset := DecodeAddress(addr)
	log.Debug().Int("tag", tag).Int("setID", setID).Int("blockoffset", blockOffset).Msg("[Cache] Reading data")

	// if in the cache, just send the value
	if isInCache {
		log.Debug().Msg("[Cache] Cache hit!")
		return line.data[blockOffset].data
	}
	log.Debug().Msg("[Cache] Cache miss...")

	// if the mot is not in cache
	// there is a reading burst
	mots := c.memory.BurstRead(addr, c.wordsNb)
	for k, mot := range mots {
		// writing each word in the cache
		c.Write(EncodeAddress(tag, setID, k), mot)
	}
	return c.memory.Read(addr)
}

// Write writes some data in a specific address in the cache
func (c *Cache) Write(addr int, data int) {
	tag, setID, blockOffset := DecodeAddress(addr)
	log.Debug().Int("tag", tag).Int("setID", setID).Int("blockoffset", blockOffset).Msg("[Cache] Writing data")

	set := c.sets[setID]
	var line Line
	for k := 0; k < len(set.lines); k++ {
		line = set.lines[k]
		if line.valid == 1 {
			// if the tag was already registered
			if line.tag == tag {
				log.Debug().Msg("[Cache] Data was already in the cache, updating..")
				line.data[blockOffset].data = data
				set.lastLineWritten = k
				return
			}
		}
	}

	// else the tag wasn't registered
	log.Debug().Msg("[Cache] The tag wasn't registered, creating new line")
	pline := NewLine(c.wordsNb, c.wordSize)
	pline.valid = 1
	pline.data[blockOffset].data = data
	pline.tag = tag

	// write on the line next to the last one or the first if empty
	idx := (set.lastLineWritten + 1) % len(set.lines)
	set.lines[idx] = *pline
}

// WriteThrough writes data on both the cache and the memory
func (c *Cache) WriteThrough(addr int, data int) {
	// first write in cache
	c.Write(addr, data)
	// then write in memory
	c.memory.Write(addr, data)
}

// GetMemory returns the pointer to the linked memory
func (c * Cache) GetMemory() *Memory {
	return c.memory
}

// SetMemory sets the new memory for the cache
func (c *Cache) SetMemory(m *Memory) {
	c.memory = m
}

// --------------- SET --------------

// Set is a structure for cache Sets data
type Set struct {
	lines []Line
	lastLineWritten int
}

// NewSet creates a new set based on the cache settings
func NewSet(linesNb int, wordsNb int, wordSize int) *Set {
	s := &Set{}
	s.lastLineWritten = -1
	for i := 0; i < linesNb; i++ {
		s.lines = append(s.lines, *NewLine(wordsNb, wordSize))
	}
	return s
}

// --------------- LINE --------------

// Line is a structure containing data for a cache line
type Line struct {
	valid 	int
	tag 	int
	data 	[]Word
	wordsNb int 
}

// NewLine creates a new line basd on the cache settings
func NewLine(nbMots int, sizeMot int) *Line {
	l := &Line{}
	for i := 0; i < nbMots; i++ {
		l.data = append(l.data, Word{size: sizeMot})
	}
	return l
}

// --------------- MOT --------------

// Word is a structure containing a cache word data
type Word struct {
	size int
	data int
}