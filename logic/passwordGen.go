package logic

import (
	"math/rand"
	"time"
)

//to generate different password each time

func init() {
	rand.Seed(time.Now().Unix())
}

func PasswordGen(length int, isSpecialLetters bool, isDigit bool) string {

	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	if isSpecialLetters{
		charset += "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	}

	if isDigit{
		charset += "1234567890"
	}

	password := make([]byte, length)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	print(string(password))

	return string(password)

}
