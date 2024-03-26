package main

import "log"

func main() {

	// first init key pair
	initCrypto()

	// get private key
	puk := getIdRsaPubStr()
	log.Printf("INFO: public key - %s", puk)

	// get private key
	prk := getIdRsaStr()
	log.Printf("INFO: private key - %s", prk)

	data := "anousone"

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
	}

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
