package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

var SERVER = ""
var PATH = ""

func randomValue(value int) {
	fmt.Scan(&value)
	for i := 0; i < value; i++ {
		randVal := rand.Intn(1000)
		fmt.Println("Random of value: ", randVal)
		fmt.Printf("Range %v of your input %v\n", i+1, value)
		fmt.Println("----------------------------")
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Need SERVER + PATH!")
		return
	}

	SERVER = arguments[1]
	PATH = arguments[2]
	fmt.Println("Connecting to:", SERVER, "at", PATH)

	randomValue(3)

	url := url.URL{Scheme: "ws", Host: SERVER, Path: PATH}
	c, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer c.Close()
}
