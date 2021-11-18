package main

import "fmt"

//START OMIT
func main() {
	theMine := []string{"rock", "ore1", "ore2", "rock", "ore3"}

	foundOre := finder(theMine)
	fmt.Println(foundOre)

	minedOre := miner(foundOre)
	fmt.Println(minedOre)

	smeltedOre := smelter(minedOre)
	fmt.Println(smeltedOre)
}

//END OMIT

func finder(mine []string) []string {
	var out []string
	for _, v := range mine {
		if v[:3] == "ore" {
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
