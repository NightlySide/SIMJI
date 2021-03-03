package cache

// Memory is a structure containing data from a memory linked
// to a cache
type Memory struct {
	size     int
	blocSize int
	data     []int
}

// NewMemory creates a new memory based on the size
func NewMemory(size int, blocSize int) *Memory {
	mem := &Memory{
		size:     size,
		blocSize: blocSize,
		data:     make([]int, size),
	}

	return mem
}

// Write writes data to a specific address
func (m *Memory) Write(addr int, data int) {
	_, setID, blockOffset := DecodeAddress(addr)
	m.data[setID * m.blocSize+ blockOffset] = data
}

// Read reads data from a specific address
// returns 0 if there is no data to be found
func (m *Memory) Read(addr int) int {
	if m.CheckData(addr) {
		_, setID, blockOffset := DecodeAddress(addr)
		return m.data[setID * m.blocSize+ blockOffset]
	}
	return 0
}

// BurstRead returns several words stored in memory
// called by the cache is the cache hits a miss
func (m *Memory) BurstRead(addr int, burstSize int) []int {
	// on lit un maximum de données pour remplir une ligne
	// complète de la mémoire
	var res []int
	_, setID, blockOffset := DecodeAddress(addr)

	// le décalage ne peut pas être plus grand
	// que le burst size
	if blockOffset >= burstSize {
		return []int{}
	}

	// iterating over the blocs
	for k := 0; k < burstSize; k++ {
		res = append(res, m.data[setID * m.blocSize + k])
	}
	return res
}

// CheckData is checking whether the address is stored in the memory
func (m *Memory) CheckData(addr int) bool {
	_, setID, blockOffset := DecodeAddress(addr)
	return len(m.data) > setID * m.blocSize + blockOffset
}

func (m *Memory) GetData() []int {
	return m.data
}

func (m *Memory) SetValueFromIndex(index int, value int) {
	m.data[index] = value
}

func (m *Memory) GetValueFromIndex(index int) int {
	return m.data[index]
}

// aide : https://github.com/michael-ross-scott/Cache-Simulator/blob/master/memoryManager.java