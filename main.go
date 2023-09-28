package main

import (
	"fmt"
	"red/src"

)

func main() {
	var m1 red.Shop
	var c1 red.Character
	var g1 red.Goblin
/* 	m1.items = map[string]int{"Hp Potion": 1, "Poison Potion": 1, "SpellBook : FireBall": 1}
	m1.price = 0
	m1.quantity = 1 */
	fmt.Print("\033[H\033[2J")
	c1.SetName()
	fmt.Print("\033[H\033[2J")
	c1.ClassChoice()
	fmt.Print("\033[H\033[2J")
	for {
		fmt.Println("-------------------------------------")
		fmt.Println("Select an option")
		red.Menu1(&c1, &m1, &g1)
	}
}

