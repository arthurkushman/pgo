package pgo

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
)

// Md5 returns simple md5 in hex generated from string s
func Md5(s string) string {
	h := md5.New()
	return hashing(h, s)
}

// Sha1 returns simple sha1 in hex generated from string s
func Sha1(s string) string {
	h := sha1.New()
	return hashing(h, s)
}

// Sha2 returns simple sha1 in hex generated from string s
func Sha2(s string) string {
	h := sha256.New()
	return hashing(h, s)
}

func hashing(h hash.Hash, s string) string {
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
