package main

import "fmt"

type Message struct {
	Text string
}

func (m Message) Print() {
	fmt.Println(m.Text)
}

func main() {
	msg := Message{Text: "hello, world!"}
	msg.Print()
}
