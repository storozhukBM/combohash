package combohash

import (
	"github.com/zeebo/xxh3"
	"hash/crc32"
)

func Hash32(bytes []byte) uint32 {
	if len(bytes) <= 64 || len(bytes) >= 512 {
		return uint32(xxh3.Hash(bytes))
	}
	return crc32.ChecksumIEEE(bytes)
}

func Hash64(bytes []byte) uint64 {
	return xxh3.Hash(bytes)
}
