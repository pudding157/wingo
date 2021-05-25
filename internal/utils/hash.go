package utils

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// get text to hash string
func HashStr(str string) string {
	// Enter a password and generate a salted hash

	toHash := []byte(str)

	hash := HashAndSalt(toHash)

	fmt.Println("Salted Hash => ", toHash)
	return hash
}

func DehashStr(hashedStr string, text string) bool {

	// Enter a password and generate a salted hash
	toDeHash := []byte(text)
	isMatch := compareHashed(hashedStr, toDeHash)

	return isMatch

}

// hash โดยใช้วิธี bcrypt
func HashAndSalt(str []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(str, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// compare between string byte and hashed string
func compareHashed(hashedStr string, plainStr []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedStr)
	err := bcrypt.CompareHashAndPassword(byteHash, plainStr)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
