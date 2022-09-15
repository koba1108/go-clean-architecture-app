package main

import (
	"fmt"
	"github.com/koba1108/go-clean-architecture-app/internals/driver"
	"log"
	"os"
)

func main() {
	log.Println("Server running...")
	driver.Serve(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
