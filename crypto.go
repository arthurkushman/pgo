package pgo

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"log"
	"strings"
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

var hashAlgos = map[string]hash.Hash{
	"md5":  md5.New(),
	"sha1": sha1.New(),
	"sha2": sha256.New(),
}

// HashFile hashs file fileName by calculating hash md5/sha1/sha2 based on it's content
func HashFile(algo, fileName string) (string, error) {
	str, err := FileGetContents(fileName)
	if err != nil {
		return "", err
	}
	input := strings.NewReader(str)

	h := hashAlgos[algo]
	if _, err := io.Copy(h, input); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
