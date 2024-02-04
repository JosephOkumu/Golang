package main

import "fmt"

type food struct {
	preptime int
}

func mealDeliveryTime(order string) int {
	menu := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}

	i, found := menu[order]

	if !found {
		return 404
	}
	return i.preptime
}

func main() {
	fmt.Println(mealDeliveryTime("burger"))
	fmt.Println(mealDeliveryTime("chips"))
	fmt.Println(mealDeliveryTime("nuggets"))
	fmt.Println(mealDeliveryTime("burger") + mealDeliveryTime("chips") + mealDeliveryTime("nuggets"))
}
