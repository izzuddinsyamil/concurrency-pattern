package main

import (
	"fmt"
	"time"
)

// START OMIT
func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second / 2)
		fmt.Println(s)
	}
}

func main() {
	go say("hello")
	say("world")
}

// END OMIT
