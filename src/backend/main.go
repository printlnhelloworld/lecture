package main

import (
	"fmt"
	"os"

	"git.hduhelp.com/hduhelper/lecture/src/backend/routers"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ip:port")
		os.Exit(1)
	}
	r := routers.SetupRouters()
	r.Run(os.Args[1])
}
