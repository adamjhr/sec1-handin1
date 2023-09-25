package main

import (
	"fmt"
	"time"

	"github.com/adamjhr/sec1handin1/node"
	"github.com/adamjhr/sec1handin1/nodes"
)

func main() {

	fmt.Println("=================================================")
	fmt.Println("Assigment part 1, Alice and Bob exchange messages")
	fmt.Println("=================================================")
	alice := node.InitNode("Alice")
	bob := node.InitNode("Bob")
	r1 := node.InitNode("r1")
	node.AddOutgoing(alice, r1)
	node.AddOutgoing(r1, bob)
	go nodes.AliceMain(alice)
	go nodes.BobMain(bob)
	go nodes.RouterMain(r1)
	time.Sleep(4 * time.Second)
	fmt.Println("============================================")
	fmt.Println("Assigment part 2, Eve finds Bobs private key")
	fmt.Println("============================================")
	eve := node.InitNode("Eve")
	node.AddOutgoing(r1, eve)
	go nodes.AliceMain(alice)
	go nodes.BobMain(bob)
	go nodes.RouterMain(r1)
	go nodes.EveMain(eve)
	time.Sleep(4 * time.Second)
	fmt.Println("=======================================================")
	fmt.Println("Assigment part 3, Weave intercepts and modifies message")
	fmt.Println("=======================================================")
	alice = node.InitNode("Alice")
	r1 = node.InitNode("r1")
	weave := node.InitNode("Weave")
	r2 := node.InitNode("r2")
	node.AddOutgoing(alice, r1)
	node.AddOutgoing(r1, weave)
	node.AddOutgoing(weave, r2)
	node.AddOutgoing(r2, bob)
	go nodes.AliceMain(alice)
	go nodes.RouterMain(r1)
	go nodes.WeaveMain(weave)
	go nodes.RouterMain(r2)
	go nodes.BobMain(bob)
	time.Sleep(4 * time.Second)
}
