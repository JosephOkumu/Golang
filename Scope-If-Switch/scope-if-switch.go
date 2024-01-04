package main

import "fmt"

func main() {
	//Here we declare success inside if and check if it is true. This value is only accessible inside the scope of if statement.

	if success := true; success {

		fmt.Println("We're rich!")
	} else {
		fmt.Println("Where did we go wrong?")
	}

	amountStolen := 50000

	//Here we declare numofThieves inside switch since we want it to be accessed only inside the switch statement.
	switch numOfThieves := 5; numOfThieves {
	case 1:
		fmt.Println("I'll take all $", amountStolen)
	case 2:
		fmt.Println("Everyone gets $", amountStolen/2)
	case 3:
		fmt.Println("Everyone gets $", amountStolen/3)
	case 4:
		fmt.Println("Everyone gets $", amountStolen/4)
	case 5:
		fmt.Println("Everyone gets $", amountStolen/5)
	default:
		fmt.Println("There's not enough to go around...")
	}
}
