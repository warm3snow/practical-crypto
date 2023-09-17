/**
 * @Author: xueyanghan
 * @File: bcrypt.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/9/16 14:15
 */

package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := []byte("123456")

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Printf("hashedPassword[%d]: %v\n", len(hashedPassword), hashedPassword)

	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return
	}
	fmt.Println("Password was correct!")
}
