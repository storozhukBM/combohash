package combohash_test

import (
	"fmt"
	"github.com/storozhukBM/combohash"
	"hash/crc32"
	"hash/maphash"
	"math/rand"
	"testing"

	"github.com/cespare/xxhash/v2"
	"github.com/zeebo/xxh3"
)

var Count = 0

func BenchmarkHash(b *testing.B) {
	for s := 4; s <= 1024; s *= 2 {
		for j := 0; j < 3; j++ {
			dif := (s*2 - s) / 3
			n := s + (dif * j)

			b.Run(fmt.Sprintf("type:crc32ieee;hash:32bit;bytes:%v", n), func(b *testing.B) {
				target := make([]byte, n+b.N)
				rand.Read(target)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					Count += int(crc32.ChecksumIEEE(target[i : i+n]))
				}
			})
			b.Run(fmt.Sprintf("type:xxhash;hash:32bit;bytes:%v", n), func(b *testing.B) {
				target := make([]byte, n+b.N)
				rand.Read(target)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					Count += int(xxhash.Sum64(target[i : i+n]))
				}
			})
			b.Run(fmt.Sprintf("type:maphash;hash:32bit;bytes:%v", n), func(b *testing.B) {
				target := make([]byte, n+b.N)
				rand.Read(target)
				b.ResetTimer()
				hash := maphash.Hash{}
				for i := 0; i < b.N; i++ {
					hash.Reset()
					_, _ = hash.Write(target[i : i+n])
					Count += int(hash.Sum64())
				}
			})
			b.Run(fmt.Sprintf("type:xxh3;hash:32bit;bytes:%v", n), func(b *testing.B) {
				target := make([]byte, n+b.N)
				rand.Read(target)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					Count += int(xxh3.Hash(target[i : i+n]))
				}
			})
			b.Run(fmt.Sprintf("type:combined;hash:32bit;bytes:%v", n), func(b *testing.B) {
				target := make([]byte, n+b.N)
				rand.Read(target)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					Count += int(combohash.Hash32(target[i : i+n]))
				}
			})

			b.Run(fmt.Sprintf("type:xxhash;hash:64bit;bytes:%v", n), func(b *testing.B) {
				target := make([]byte, n+b.N)
				rand.Read(target)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					Count += int(xxhash.Sum64(target[i : i+n]))
				}
			})
			b.Run(fmt.Sprintf("type:maphash;hash:64bit;bytes:%v", n), func(b *testing.B) {
				target := make([]byte, n+b.N)
				rand.Read(target)
				b.ResetTimer()
				hash := maphash.Hash{}
				for i := 0; i < b.N; i++ {
					hash.Reset()
					_, _ = hash.Write(target[i : i+n])
					Count += int(hash.Sum64())
				}
			})
			b.Run(fmt.Sprintf("type:xxh3;hash:64bit;bytes:%v", n), func(b *testing.B) {
				target := make([]byte, n+b.N)
				rand.Read(target)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					Count += int(xxh3.Hash(target[i : i+n]))
				}
			})
			b.Run(fmt.Sprintf("type:combined;hash:64bit;bytes:%v", n), func(b *testing.B) {
				target := make([]byte, n+b.N)
				rand.Read(target)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					Count += int(combohash.Hash64(target[i : i+n]))
				}
			})
		}
	}
}
