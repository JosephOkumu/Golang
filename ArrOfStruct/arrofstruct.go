package main

import "fmt"

// creating a struct of cake

type Cake struct {
	typeofCake string
	weight     int
}

func main() {
	//creating an array of the cake instance
	cakes := []Cake{{"Chocolate", 1000}, {"Carrot", 500}, {"Ice Cream", 2000}}

	//The code below prints the weight of chocolate cake
	fmt.Println(cakes[0].weight)

	//This code changes the weight of carrot cake
	cakes[1].weight = 1500
	fmt.Println(cakes)

}
