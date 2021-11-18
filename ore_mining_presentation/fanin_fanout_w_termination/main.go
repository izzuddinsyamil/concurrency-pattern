package main

import (
	"fmt"
	"sync"
)

func main() {

	done := make(chan bool)
	defer close(done)

	theMine := []string{"rock", "ore1", "ore2", "rock", "ore3"}

	oreChan := finder(done, theMine)

	m1 := miner(done, oreChan)
	m2 := miner(done, oreChan)
	m3 := miner(done, oreChan)

	smelted := smelter(done, m1, m2, m3)
	fmt.Printf("%s is smelted\n", <-smelted)
	fmt.Printf("%s is smelted\n", <-smelted)
}

func finder(done chan bool, mine []string) <-chan string {
	out := make(chan string)
	go func() {
		for _, item := range mine {
			if item[:3] == "ore" {
				select {
				case out <- item:
				case <-done:
					return
				}

			}
		}
		close(out)
	}()
	return out
}

func miner(done chan bool, ore <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for o := range ore {
			select {
			case out <- fmt.Sprintf("mined %s", o):
			case <-done:
				return
			}

		}
		close(out)
	}()
	return out
}

func smelter(done chan bool, minedOre ...<-chan string) <-chan string {
	wg := sync.WaitGroup{}
	wg.Add(len(minedOre))

	out := make(chan string)
	for _, c := range minedOre {
		go func(ore <-chan string) {
			for o := range ore {
				select {
				case out <- o:
				case <-done:
					return
				}

			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
