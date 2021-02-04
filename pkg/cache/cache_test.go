package cache

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	// TODO : implement testing
	cache := NewCache(4, 1, 2, 4)

	address := EncodeAddress(0, 0, 0)
	address2 := EncodeAddress(10, 2, 2)

	memaddr2 := EncodeAddress(10, 2, 3)
	cache.GetMemory().Write(memaddr2, 255)

	// Writethrough
	fmt.Printf("WRITETHROUGH - Adresse : 0x%08X --> 0xDEAD\n", address)
	cache.WriteThrough(address, 0xdead)
	inCache, _ := cache.InCache(address)
	fmt.Printf("Is address in cache ? : %t\n", inCache)
	fmt.Printf("READ - From cache : 0x%08X --> 0x%08X\n", address, cache.Read(address))

	// Test if value in cache
	inCache2, _ := cache.InCache(address2)
	fmt.Printf("Is address2 in cache ? : %t\n", inCache2)

	// Write but not through
	fmt.Printf("WRITE - Adresse : 0x%08X --> 0xBEEF\n", address2)
	// printing the cache
	fmt.Printf("Cache -> %+v\n", cache.sets[2].lines[0])
	cache.GetMemory().Write(address2, 0xbeef)
	// Getting the value from memory
	value := cache.Read(address2)
	fmt.Printf("READ - Not in cache : 0x%08X -> 0x%08X\n", address2, value)
	// printing the cache
	fmt.Printf("Cache -> %+v\n", cache.sets[2].lines[0])

	//printing the memory
	fmt.Printf("Memory -> %v\n", cache.GetMemory().GetData())
}