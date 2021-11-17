package main

import (
	"fmt"
	"sync"
)

func main() {
	theMine := []string{"rock", "ore1", "ore2", "rock", "ore3"}
	oreChan := make(chan string)
	minedOreChan := make(chan string)

	wg := &sync.WaitGroup{}
	wg.Add(3)

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

	// Ore Breaker
	go func() {
		for foundOre := range oreChan {
			fmt.Println("From Finder: ", foundOre)
			minedOreChan <- "mined " + foundOre
		}
		close(minedOreChan)
		wg.Done()
	}()

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
