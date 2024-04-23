package main

import (
	"fmt"
	db "rest_api/config"
)

func main() {
	fmt.Println("Connecting to DB...")
	db.Connection()

}
