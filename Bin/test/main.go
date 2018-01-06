package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {

	for i := 0; i < 10; i++ {
		salt := make([]byte, 32)
		_, err := io.ReadFull(rand.Reader, salt)
		if err != nil {
			fmt.Println("err: ", err)
		}
		fmt.Println(salt)
	}
}
