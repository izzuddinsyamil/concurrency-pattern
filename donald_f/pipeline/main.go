package main

import "fmt"

type item struct {
	price    int
	category string
}

func main() {
	c := gen(
		item{8, "shirt"},
		item{20, "shoe"},
		item{24, "shoe"},
		item{4, "drink"},
	)

	out := discount(c)
	for processed := range out {
		fmt.Println("Category:", processed.category, "Price:", processed.price)
	}
}

func gen(items ...item) <-chan item {
	out := make(chan item, len(items))
	for _, i := range items {
		out <- i
	}

	close(out)
	return out
}

func discount(items <-chan item) <-chan item {
	out := make(chan item)
	go func() {
		defer close(out)
		for i := range items {
			if i.category == "shoe" {
				i.price /= 2
			}
			out <- i
		}
	}()
	return out
}
