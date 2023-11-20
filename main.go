package main

import (
	"fmt"
	"main.go/u"
	"time"
)

func main() {

	fmt.Printf("Version %s\n", u.V())
	fmt.Printf("OK!Success! uuid = %s", u.GetUid())
	for {
		time.Sleep(1)
	}
}
