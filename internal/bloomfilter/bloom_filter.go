package bloomfilter

import (
	"hash"
	"log"
	"strconv"
	"time"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	filter []bool
	size   int
}

var bloomfiler *BloomFilter

var mHasher hash.Hash32

func init() {
	mHasher = murmur3.New32WithSeed(uint32(time.Now().UnixNano()))
}

func murmurhash(key string, size int) int {
	mHasher.Write([]byte(key))
	result := mHasher.Sum32() % uint32(size)
	mHasher.Reset()
	return int(result)
}

func NewBloomFilter(size int) *BloomFilter {
	return &BloomFilter{
		filter: make([]bool, size),
		size:   size,
	}
}

func Initialize(n string) {
	size, err := strconv.Atoi(n)
	if err != nil {
		log.Println("Invalid size input:", err)
		return
	}
	bloomfiler = NewBloomFilter(size)
}

func (b *BloomFilter) Add(key string) {
	idx := murmurhash(key, b.size)
	b.filter[idx] = true
}

func (b *BloomFilter) Exists(key string) bool {
	idx := murmurhash(key, b.size)
	return b.filter[idx]
}

func GetBloomFilter() *BloomFilter {
	return bloomfiler
}
