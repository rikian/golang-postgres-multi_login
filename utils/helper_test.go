package utils

import (
	"log"
	"testing"
)

var testToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJpa2lhbmZhaXNhbEBnbWFpbC5jb20iLCJleHAiOjE2NTMzMzU2NzYsInBhc3N3b3JkIjoicjRoNDUxNC4uLiJ9.37GFs1iU_iVhpowikI3MmxbEHE7_4lWpSX7awM-97OM"

func TestEncryptToken(t *testing.T) {
	encryptToken, err := EncryptToken(60)
	if err {
		log.Println("failed")
		return
	} 

	log.Println(encryptToken)
}

func TestValidDecryptToken(t *testing.T) {
	encryptToken, _ := EncryptToken(60)
	decryptToken, err := DecryptToken(encryptToken)
	if err {
		log.Println("token expired")
		return
	}

	log.Println(decryptToken)
}

func TestExpairedDecryptToken(t *testing.T) {
	decryptToken, err := DecryptToken(testToken)
	if err {
		log.Println("token expired")
		return
	}

	log.Println(decryptToken)
}

func BenchmarkEncrypttoken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncryptToken(60)
	}
}

func BenchmarkDecryptToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DecryptToken(testToken)
	}
}