package main

import (
	//"fmt"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
)

func randomValue(c net.Conn) {
	randVal := rand.Intn(1000)
	log.Println("Server returns a random of value(1-1000) --> Your value:", randVal)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		log.Println("Please provide a port number")
		return
	} else {
		log.Println("Connection to port -->", arguments[1])
	}

	port := ":" + arguments[1]

	l, err := net.Listen("tcp4", port)
	assertNotNill(err)
	defer l.Close()

	c, err := l.Accept()
	assertNotNill(err)
	randomValue(c)
}

func assertNotNill(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
