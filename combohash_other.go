//go:build !arm64
// +build !arm64

package combohash

func Hash32(bytes []byte) uint32 {
	return uint32(xxh3.Hash(bytes))
}

func Hash64(bytes []byte) uint32 {
	return xxh3.Hash(bytes)
}
