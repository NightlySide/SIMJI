package vm

type Memory struct {
	size int
	data []int
}

func NewMemory(size int) *Memory {
	return &Memory{
		size: size}
}

func (m *Memory) Write(addr int, data int) {

}

func (m *Memory) Read(addr int) int {
	return 0
}

func (m *Memory) BurstRead(addr int, burstSize int) {

}

func (m *Memory) CheckAddress(addr int, size int) {

}

func (m *Memory) CheckData(addr int, nbits int) {

}



type Bloc struct {
	valid int
	tagBits int
	size int
}

func NewBloc() *Bloc {
	return &Bloc{size: 16}
}

// aide : https://github.com/michael-ross-scott/Cache-Simulator/blob/master/memoryManager.java