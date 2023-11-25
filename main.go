package main

import (
	"fmt"
	"main.go/u/orders"
	"time"
)

func main() {
	fmt.Printf("Version %s", orders.Version)
	for {
		time.Sleep(1)
	}
}
