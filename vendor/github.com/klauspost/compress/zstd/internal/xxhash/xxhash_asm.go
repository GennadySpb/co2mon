<<<<<<<< HEAD:vendor/github.com/klauspost/compress/zstd/internal/xxhash/xxhash_asm.go
//go:build (amd64 || arm64) && !appengine && gc && !purego && !noasm
========
//go:build (amd64 || arm64) && !appengine && gc && !purego
>>>>>>>> f62e69b (Update go modules and vendor dir):vendor/github.com/cespare/xxhash/v2/xxhash_asm.go
// +build amd64 arm64
// +build !appengine
// +build gc
// +build !purego
// +build !noasm

package xxhash

// Sum64 computes the 64-bit xxHash digest of b.
//
//go:noescape
func Sum64(b []byte) uint64

//go:noescape
func writeBlocks(s *Digest, b []byte) int
