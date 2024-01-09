package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//type RegisterRequest struct {
//	UserName string
//	Password string
//}

var RegisteredUsers map[string]string = make(map[string]string)

//var registeredUsers RegisteredUsers = RegisteredUsers{}

func HashPassword(password string) (hashedPassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPassword(hashedPassword, password string) (isPasswordValid bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}

	return true
}

func Login(userName, password string) string {
	if CheckPassword(RegisteredUsers[userName], password) {

		_, str, err := token.Encode(map[string]interface{}{"username": userName})
		if err != nil {
			fmt.Println("token error occred:", err)
		}

		fmt.Println("token assigned:", str)

		fmt.Println("user " + userName + " authenticated")
		return str

	}
	fmt.Println("access denied")
	return ""
}

func Register(userName, password string) string {
	hash, err := HashPassword(password)
	if err != nil {
		return "error hapened:" + err.Error()
	}

	RegisteredUsers[userName] = hash

	fmt.Println("user, password, hash:", userName, password, hash)
	return "user registerred"
}
