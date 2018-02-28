package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
)

func main() {
	flag.Parse()
	secret := os.Args[1]
	code, _ := totp.GenerateCode(secret, time.Now())
	fmt.Println(code)
}
