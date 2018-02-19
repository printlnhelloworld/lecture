package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ip:port")
		os.Exit(1)
	}
	r := setupRouters()
	r.Run(os.Args[1])
}
