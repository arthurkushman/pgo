package pgo

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

// Md5 returns simple md5 in hex generated from string s
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Sha1 returns simple sha1 in hex generated from string s
func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
