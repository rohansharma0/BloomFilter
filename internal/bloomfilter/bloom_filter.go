package bloomfilter

import (
	"hash"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/rohansharma0/bloomfiler/pkg/redisclient"
	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	filter []bool
	size   int
	mutex  sync.RWMutex
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
	log.Println("Bloom filter initialized with size:", size)
}

func (b *BloomFilter) Add(key string) {
	idx := murmurhash(key, b.size)
	b.mutex.Lock()
	b.filter[idx] = true
	redisclient.Client.Set(redisclient.Ctx, key, true, 5*24*time.Hour).Err()
	b.mutex.Unlock()
}

func (b *BloomFilter) Exists(key string) bool {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	idx := murmurhash(key, b.size)
	log.Println(b.filter[idx])
	return b.filter[idx]
}

func GetBloomFilter() *BloomFilter {
	if bloomfiler == nil {
		log.Println("WARNING: Bloom filter not initialized")
	}
	return bloomfiler
}
