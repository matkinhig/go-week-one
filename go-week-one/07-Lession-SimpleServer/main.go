package main

import (
	"fmt"
	"net/http"
	"github.com/matkinhig/07-lession-simpleserver/handler"
)

func main() {
	fmt.Println("start golang")

	http.HandleFunc("/api/v1/health", handler.)
}
