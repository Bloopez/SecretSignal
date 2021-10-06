package main 

import (
	"fmt"
	//"context"
	//"time"
    "crypto/aes"
    "crypto/cipher"
	"crypto/rand"
	"io"
	"encoding/hex"
	"encoding/json"
)


var nonce, _ = hex.DecodeString("29e11e1abbd52e10c5b537f6")
var key, _ = hex.DecodeString("b7df6e9682d3caeace0216e133e6eb533e0054e5bfbebdbe5eb1d59d6d266722")

type cmd_obj struct {
    CmdType string
    CmdContent string
	ClientName string
}
//generate keys
func gen_info() {
    key := make([]byte, 32)
    if _, err := io.ReadFull(rand.Reader, key); err != nil {
      panic(err.Error())
	}

    nonce := make([]byte, 12)
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
      panic(err.Error())
    }

	fmt.Printf("Nonce:\t%x", nonce)
    fmt.Printf("\nKey:\t%x\n", key)
}
// encrypt text
func enc_info(plaintext string) {
	block, err := aes.NewCipher(key)
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(plaintext), nil)
	fmt.Printf("%x\n", ciphertext)
}
// decrypt payload
func dec_info(ciphertext string) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	decode, _ := hex.DecodeString(ciphertext)

	plaintext, err := aesgcm.Open(nil, nonce, decode, nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s\n", plaintext)
}
// specify command to be executed in CmdContent
func main() {
	m := cmd_obj{
		CmdType: "ps",
		CmdContent: "whoami",
		ClientName: "test",
	}
	res1B, _ := json.Marshal(m)
	enc_info(string(res1B))
	dec_info("11b2070350ef228c2408fe790ae2225089ec4152465a54ec54e7c2cb5c7d08523de9d6330e875bc49f5fc6e172548a8a57dd8408dc99bb0849ca82bed26f186b14af77b037d3162c9be2")
}