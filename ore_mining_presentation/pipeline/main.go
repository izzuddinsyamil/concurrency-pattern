package main

import (
	"fmt"
	"sync"
)

// START1 OMIT
func main() {
	theMine := []string{"rock", "ore1", "ore2", "rock", "ore3"}
	oreChan := make(chan string)      // HL
	minedOreChan := make(chan string) // HL

	wg := &sync.WaitGroup{}
	wg.Add(3)
	// END1 OMIT

	// START2 OMIT
	// Finder
	go func(mine []string) {
		for _, item := range mine {
			if item[:3] == "ore" {
				oreChan <- item
			}
		}
		close(oreChan)
		wg.Done()
	}(theMine)
	// END2 OMIT

	// START3 OMIT
	// Ore Breaker
	go func() {
		for foundOre := range oreChan {
			fmt.Println("From Finder: ", foundOre)
			minedOreChan <- "mined " + foundOre
		}
		close(minedOreChan)
		wg.Done()
	}()
	// END3 OMIT

	// START4 OMIT
	// Smelter
	go func() {
		for minedOre := range minedOreChan {
			fmt.Println("From Miner: ", minedOre)
			fmt.Printf("From Smelter: %s is smelted\n", minedOre)
		}
		wg.Done()
	}()

	wg.Wait()
}

// END4 OMIT
