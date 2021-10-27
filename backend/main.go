package main

import (
	"flag"
	"server/chat"
)

var (
	host = flag.String("host", "http://localhost", "set host")
	port = flag.String("port", ":8000", "set port")
)

func main() {
	chat.Start(*host, *port)
}
