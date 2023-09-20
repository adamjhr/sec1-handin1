package person

type person struct {
	name     string
	recvChan chan message
}
