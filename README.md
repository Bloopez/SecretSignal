# DNS over HTTPS backdoor

This program grabs the DNS TXT record from a specified domain to obtain and execute any command given.

## Features
AES Encryption - Command(payload) is encrypted in AES for delivery to the DNS TXT record
SSL Encryption - Retrieval of payload goes through HTTPS
Execute recieved command via powershell

### Files & Functions

#### generator.go
Responsible for generating the payload


#### main.go
Responsible for retrieving and executing the payload from DNS TXT record



## How to build/use

>**_NOTE:_** So far this only works with windows. 
To use, replace the domain in the main.go file with your domain and run 
`GOOS=windows GOARCH=amd64 go build main.go`
to compile the executable


