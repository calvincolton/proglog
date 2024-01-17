package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/calvincolton/proglog/internal/server"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "port number")
	flag.Parse()

	addr := fmt.Sprintf(":%s", port)
	srv := server.NewHTTPServer(addr)
	log.Fatal(srv.ListenAndServe())
}
