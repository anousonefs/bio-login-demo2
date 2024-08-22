package main

import (
	"fmt"
	"log"
)

func main() {

	// first init key pair
	/* initCrypto() */

	publicKey := getIdRsaPubStr()
	log.Printf("INFO: public key - %s", publicKey)

	privateKey := getIdRsaStr()
	log.Printf("INFO: private key - %s", privateKey)

	data := "anousone@gmail.com.dfbda-da66-4a5a-9f09-56a7979aaab5.1711434369"

	// sign
	sig, err := sign(data, getIdRsa())
	if err != nil {
		log.Printf("ERROR: fail to sign - %s", err.Error())
	}
	log.Printf("INFO: signature - %s", sig)

	// verify
	err = verify(data, sig, getIdRsaPub())
	if err != nil {
		log.Printf("ERROR: fail to verify - %s", err.Error())
	} else {
		fmt.Printf("verify success\n")
	}

	// encryp - decrypt

	// encrypt
	ciperText, err := encrypt(data, getIdRsaPub())
	if err != nil {
		log.Printf("ERROR: fail encrypt - %s", err.Error())
	}
	log.Printf("INFO: ciper text - %s", ciperText)

	// decrypt
	plainText, err := decrypt(ciperText, getIdRsa())
	if err != nil {
		log.Printf("ERROR: fail decrypt - %s", err.Error())
	}
	log.Printf("INFO: plain text - %s", plainText)
}
