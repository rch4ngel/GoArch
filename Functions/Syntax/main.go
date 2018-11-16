package main

import "fmt"

func main() {
	block("Ryu")
	hit("Ken")

	d, o := finish("Ken")
	fmt.Println("Ryu you finished:", o, "with:", d, "damage!")
}

func hit(w string) (d int) {
	fmt.Println("You just hit:", w)

	return 20
}

func block(c string) {
	fmt.Println(c, "! You blocked their hit!")
}

func finish(c string) (d int, o string) {
	fmt.Println(c, "!", "You destroyed your opponent")

	return 43, "Ken"
}
