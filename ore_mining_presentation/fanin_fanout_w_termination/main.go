package main

import (
	"fmt"
	"sync"
)

// START1 OMIT
func main() {
	done := make(chan bool) // HL
	defer close(done)       // HL

	theMine := []string{"rock", "ore1", "ore2", "rock", "ore3"}

	// END1 OMIT
	oreChan := finder(done, theMine)

	m1 := miner(done, oreChan)
	m2 := miner(done, oreChan)
	m3 := miner(done, oreChan)

	smelted := smelter(done, m1, m2, m3)
	fmt.Printf("%s is smelted\n", <-smelted)
	fmt.Printf("%s is smelted\n", <-smelted)
}

// START2 OMIT
func finder(done chan bool, mine []string) <-chan string {
	out := make(chan string)
	go func() {
		for _, item := range mine {
			if item[:3] == "ore" {
				select { // HL
				case out <- item: // HL
				case <-done: // HL
					return // HL
				} // HL
			}
		}
		close(out)
	}()
	return out
}

// END2 OMIT

// START3 OMIT
func miner(done chan bool, ore <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for o := range ore {
			select { // HL
			case out <- fmt.Sprintf("mined %s", o): // HL
			case <-done: // HL
				return // HL
			} // HL
		}
		close(out)
	}()
	return out
}

// END3 OMIT

// START4 OMIT
func smelter(done chan bool, minedOre ...<-chan string) <-chan string {
	wg := sync.WaitGroup{}
	wg.Add(len(minedOre))

	out := make(chan string)
	for _, c := range minedOre {
		go func(ore <-chan string) {
			for o := range ore { // HL
				select { // HL
				case out <- o: // HL
				case <-done: // HL
					return // HL
				} // HL
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

// END4 OMIT
