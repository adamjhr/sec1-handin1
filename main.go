package main

import (
	"fmt"
	"github.com/adamjhr/sec1-handin1/message"
)

func main() {

	fmt.Println("Hello")
	m := Message{Sender: "test", Value: 10, Pub_key: 20}
	fmt.Println(m.Sender, m.Value, m.Pub_key)

}
