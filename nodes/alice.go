package nodes

import (
	"fmt"
	"math/big"
	"math/rand"
)

func aliceMain(self *Node) {
	g := 666
	p := 6661
	pub_key := 2227
	m := 2000
	y := rand.Intn(p)
	fmt.Println("I'm encrypting with y = ", y)
	// Encrypt message
	m_encr_exp := new(big.Int).Exp(big.NewInt(int64(pub_key)), big.NewInt(int64(y)), nil)
	m_encr_mul := new(big.Int).Mul(m_encr_exp, big.NewInt(int64(m)))
	m_encr_big := new(big.Int).Mod(m_encr_mul, big.NewInt(int64(p)))
	m_encr := m_encr_big.Uint64()
	fmt.Println("The encrypted message is ", m_encr)
	// Calculate key
	key_big := new(big.Int).Exp(big.NewInt(int64(g)), big.NewInt(int64(y)), big.NewInt(int64(p)))
	key := key_big.Uint64()
	fmt.Println("My own public key is ", key)

	message := Message{Sender: "Alice", Value: m_encr, Pub_key: key}

	for _, channel := range &self.outgoing {
		&channel <- message
	}
}
