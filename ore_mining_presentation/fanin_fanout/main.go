package main

import (
	"fmt"
	"sync"
)

// START4 OMIT
func main() {
	theMine := []string{"rock", "ore1", "ore2", "rock", "ore3"}

	oreChan := finder(theMine)

	m1 := miner(oreChan)
	m2 := miner(oreChan)
	m3 := miner(oreChan)

	for smelted := range smelter(m1, m2, m3) {
		fmt.Printf("%s is smelted\n", smelted)
	}
}

// END4 OMIT

// START1 OMIT
func finder(mine []string) <-chan string {
	out := make(chan string)
	go func() {
		for _, item := range mine {
			if item[:3] == "ore" {
				out <- item
			}
		}
		close(out)
	}()
	return out
}

// END1 OMIT

// START2 OMIT
func miner(ore <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for o := range ore {
			out <- fmt.Sprintf("mined %s", o)
		}
		close(out)
	}()
	return out
}

// END2 OMIT

// START3 OMIT
func smelter(minedOre ...<-chan string) <-chan string {
	wg := sync.WaitGroup{}
	wg.Add(len(minedOre))

	out := make(chan string)
	for _, c := range minedOre {
		go func(ore <-chan string) {
			for o := range ore {
				out <- o
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

// END3 OMIT
