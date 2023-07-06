/**
 * @Author: xueyanghan
 * @File: main.go
 * @Version: 1.0.0
 * @Description: TODO
 * @Date: 2023/3/13 16:17
 */
package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := []byte("123456")

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(hashedPassword))
	fmt.Println(string(hashedPassword))

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println(err) // nil means it is a match
}
