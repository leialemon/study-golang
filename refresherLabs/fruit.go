package main

import "fmt"

type Item struct {
	Name  string
	Price float64
}

func getItemsInPriceRange(items []Item, minPrice, maxPrice float64) []Item {
	var budget []Item
	for i:=0;i<len(items);i++{
		element := items[i]
		if element.Price >= minPrice && element.Price <= maxPrice {
			budget = append(budget, element)
		}
	}
	return budget
}