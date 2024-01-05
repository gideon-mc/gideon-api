package auth

import (
	"log"
	"math"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Printf("(bcrypt) %s", err)
	}
	return string(bytes)
}

func CompareHashAndPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsPasswordSecure(password string) bool {
	return GetShannonEntropy(password) > 40
}

func GetShannonEntropy(value string) int {
	frequency := make(map[rune]float64)

	for _, i := range value {
		frequency[i]++
	}

	var sum float64

	for _, v := range frequency {
		f := v / float64(len(value))
		sum += f * math.Log2(f)
	}

	return int(math.Ceil(sum*-1)) * len(value)
}
