package main

import "fmt"

func main() {
	theMine := []string{"rock", "ore", "ore", "rock", "ore"}

	foundOre := finder(theMine)
	fmt.Println(foundOre)

	minedOre := miner(foundOre)
	fmt.Println(minedOre)

	smeltedOre := smelter(minedOre)
	fmt.Println(smeltedOre)
}

func finder(mine []string) []string {
	var out []string
	for _, v := range mine {
		if v == "ore" {
			out = append(out, v)
		}
	}

	return out
}

func miner(ores []string) []string {
	var out []string
	for i := 0; i < len(ores); i++ {
		out = append(out, "minedOre")
	}

	return out
}

func smelter(ores []string) []string {
	var out []string
	for i := 0; i < len(ores); i++ {
		out = append(out, "smeltedOre")
	}

	return out
}
