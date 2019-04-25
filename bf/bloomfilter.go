package bf

import (
	"github.com/spaolacci/murmur3"
	"github.com/willf/bitset"
)

// todo allow variable number of hash functions

type bloomFilter struct {
	size uint
	bucket *bitset.BitSet
}

// NewBlommFilter return an empty bloom filter with the size n
func NewBloomFilter(n uint) *bloomFilter {
	return &bloomFilter{n,bitset.New(n)}
}

// hashValues return the positions to set 1 in the bloom filter bucket
func hashValues(data []byte) [4]uint64{
	h := murmur3.New128()
	h.Write(data)
	h1,h2 := h.Sum128()
	h.Write([]byte{1})
	h3,h4 := h.Sum128()
	return [4]uint64{h1,h2,h3,h4}
}

// Insert elements into the bloom filter
func (bf *bloomFilter) Insert(data []byte)  {
	hv := hashValues(data)
	for _,v := range hv {
		bf.bucket.Set(uint(v)%bf.size)
	}
}

// Containts check if the bloom filter contains the value
func (bf *bloomFilter) Contains(data []byte) bool{
	hv := hashValues(data)
	for _,v := range hv {
		if bf.bucket.Test(uint(v)%bf.size) {
			return true
		}
	}
	return false
}