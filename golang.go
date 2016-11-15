package main

import "fmt"

const N = 22
const recipient = 17
type Token struct {
    data string
    recipient int
}

func link(id int, this <-chan Token, next chan<- Token) {
	for {
		msg := <- this
		if id == msg.recipient {
			fmt.Println(msg.data, id)
        	} else {
                	fmt.Println("next, I am",id)
            		next <- msg
        	}
	}
}

func main() {
	start := make(chan Token)
	this := start
	var next chan Token
	for i := 0; i < N; i++ {
		next = make(chan Token)
		go link(i, this, next)
		this = next
	}
	start <- Token{"Hello it's me!", recipient}
	var input string
    	fmt.Scanln(&input)
}
