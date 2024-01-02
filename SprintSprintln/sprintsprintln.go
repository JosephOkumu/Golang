package main

import "fmt"

func main() {
	step1 := "Breathe in..."
	step2 := "Breathe out..."

	// Add your code below:
	//The sprintln just works like the println method, but it doesn't print. It returns
	//The sprint just works like the print method but it doesn't print, it returns.
	//The comma for arguments in print does not print with space. Same to sprint
	meditation := fmt.Sprintln(step1, step2)
	fmt.Println(meditation)

	//fmt.Sprintf() works very similarly to fmt.Printf(), the major difference is that fmt.Sprintf()
	// returns its value instead of printing it out!

	template := "I wish I had a %v."
	pet := "puppy"

	var wish string

	// Add your code below:
	wish = fmt.Sprintf(template, pet)

	fmt.Println(wish)

}
