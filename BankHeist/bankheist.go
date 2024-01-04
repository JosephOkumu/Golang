package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//The line 11 below is not necessary.
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)
	//fmt.Println(eludedGuards)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)
	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("The vault can't be opened.")
	}
	leftSafely := rand.Intn(5)
	if isHeistOn == true {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Bank Heist Failed!!!")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		case 2:
			isHeistOn = false
			fmt.Println("CCTV cameras are mounted almost on every wall.")
		case 3:
			isHeistOn = false
			fmt.Println("The alarm went off.")
		default:
			fmt.Println("Start the getaway car!")
		}
		if isHeistOn {
			//This line below approximates the amount present in the bank.
			amtStolen := 10000 + rand.Intn(1000000)
			fmt.Println("The amount stolen is:", amtStolen)
		}

	}
	fmt.Printf("%v", isHeistOn)
}
