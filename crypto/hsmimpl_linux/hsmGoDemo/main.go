package main

import (
	"github.com/warm3snow/practical-crypto/crypto/hsmimpl_linux/hsmGoDemo/swsds"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Run")

	// swsds.Test_Encrypt_Decrypt()
	swsds.Test_Encrypt_Decrypt_2()

	log.Println("End")
}
