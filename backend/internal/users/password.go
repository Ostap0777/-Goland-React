package users

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)


const (
	pbkdf2Algoritm = "pbkdf2_sha256"

	defaultIterations = 600_000

	saltBytes = 16
	keyBytes = 32
)

var dummySalt = []byte("timing-equalizer")


func HashPassword(password string) (string, error) {
	return hashPassword(password, defaultIterations)
}

func hashPassword(password string, iterations int) (string, error) {
	salt := make([]byte, saltBytes)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("generate salt: %w", err)
	}
	key := pbkdf2Key([]byte(password), salt, iterations, keyBytes)
	return fmt.Sprintf("%s$%d$%s$%s",
	pbkdf2Algoritm,
	iterations,
	base64.RawStdEncoding.EncodeToString(salt),
	base64.RawStdEncoding.EncodeToString(key),
), nil
}


func VerifyPassword(encoded, password string) bool {
	parts := strings.Split(encoded, "$")
	if len(parts) != 4 || parts[0] != pbkdf2Algoritm {
		return false
	}
	iteration, err := strconv.Atoi(parts[1])
	if err != nil || iteration < 1 {
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[2]) 
	if err != nil || len(salt) == 0  {
		return false
	}

	want, err := base64.RawStdEncoding.DecodeString(parts[3]) 
	if err != nil || len(want) == 0 {
		return false
	}

	got := pbkdf2Key([]byte(password), salt, iteration, len(want))
	return subtle.ConstantTimeCompare(got,want) == 1
}

func pbkdf2Key(password, salt []byte, iterations, keyLen int) []byte {
	if keyLen <= 0 {
		return nil
	}

	prf := hmac.New(sha256.New, password)
	hashLen := prf.Size()
	numBlocks := (keyLen + hashLen - 1) / hashLen


	var blockNum [4]byte
	dk := make([]byte, 0, numBlocks*hashLen)
	u := make([]byte, 0, hashLen)
	for block := 1; block <= numBlocks; block++ {
		prf.Reset()
		prf.Write(salt)
      blockNum[0] = byte(block >> 24)
      blockNum[1] = byte(block >> 16)
      blockNum[2] = byte(block >> 8)
      blockNum[3] = byte(block)
		prf.Write(blockNum[:])


			t := prf.Sum(nil)
			u = append(u[:0], t...)
			for n := 2; n <= iterations; n++ {
				prf.Reset()
		      prf.Write(u)
		      u = prf.Sum(u[:0])
		      for i := range t {
			       t[i] ^= u[i]
		      }
			}
			dk = append(dk, t...)
	}
	return dk[:keyLen]
}