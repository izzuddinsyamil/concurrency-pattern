package main

import (
	"fmt"
	"sync"
	"time"
)

type item struct {
	price    int
	category string
}

func main() {
	done := make(chan bool)
	defer close(done)

	c := gen(
		item{8, "shirt"},
		item{20, "shoe"},
		item{24, "shoe"},
		item{4, "drink"},
	)

	c1 := discount(done, c)
	c2 := discount(done, c)

	out := fanIn(done, c1, c2)

	out1 := <-out
	fmt.Println("Category:", out1.category, "Price:", out1.price)

	out2 := <-out
	fmt.Println("Category:", out2.category, "Price:", out2.price)
}

func gen(items ...item) <-chan item {
	out := make(chan item, len(items))
	for _, i := range items {
		out <- i
	}

	close(out)
	return out
}

func discount(done <-chan bool, items <-chan item) <-chan item {
	out := make(chan item)
	go func() {
		defer close(out)
		for i := range items {
			time.Sleep(time.Second / 2)
			if i.category == "shoe" {
				i.price /= 2
			}
			select {
			case out <- i:
			case <-done:
				return
			}

		}
	}()
	return out
}

func fanIn(done <-chan bool, channels ...<-chan item) <-chan item {
	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	out := make(chan item)

	for _, c := range channels {
		go func(c <-chan item) {
			defer wg.Done()
			for i := range c {
				select {
				case out <- i:
				case <-done:
					return
				}
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
