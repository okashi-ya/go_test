package main

import (
	"fmt"
	"main.go/u"
	"time"
)

var Version = "test"

func main() {
	// info, ok := debug.ReadBuildInfo()
	//if ok {
	//	fmt.Printf(info.Main.Version)
	//}
	fmt.Printf("Version %s", Version)
	fmt.Printf("OK!Success! uuid = %s", u.GetUid())
	for {
		time.Sleep(1)
	}
}
