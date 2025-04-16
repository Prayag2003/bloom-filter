package bloom

import (
	"hash/fnv"
)

type Filter interface {
	Add(item string)
	Check(item string) bool
}

type BloomFilter struct {
	bitset []bool
	k      int
	size   uint32
}

func New(size uint32, k int) *BloomFilter {
	return &BloomFilter{
		bitset: make([]bool, size),
		k:      k,
		size:   size,
	}
}

func (bf *BloomFilter) Add(item string) {
	for _, h := range bf.getHashes(item) {
		bf.bitset[h%bf.size] = true
	}
}

func (bf *BloomFilter) Check(item string) bool {
	for _, h := range bf.getHashes(item) {
		if !bf.bitset[h%bf.size] {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) getHashes(item string) []uint32 {
	hashes := make([]uint32, bf.k)
	for i := 0; i < bf.k; i++ {
		h := fnv.New32a()
		h.Write([]byte(item))
		h.Write([]byte{byte(i)})
		hashes[i] = h.Sum32()
	}
	return hashes
}
