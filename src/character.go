package red

import "fmt"

type Character struct {
	name              string
	class             string
	level             int
	health_max        int
	current_health    int
	inventory         map[string]int
	skills            []string
	golds             int
	equipment         Equipment
	inventoryCapacity int
	inventoryUpgrades int
	initiative        int
}

var nickname string

func main() {
	var m1 Shop

	m1.items = map[string]int{"Hp Potion": 1, "Poison Potion": 1, "SpellBook : FireBall": 1}
	m1.price = 0
	m1.quantity = 1
}

func (c *Character) Init(name string, class string, level int, health_max int, current_health int, inventory map[string]int, skills []string, golds int) {
	c.name = name
	c.class = class
	c.level = level
	c.health_max = health_max
	c.current_health = current_health
	c.inventory = inventory
	c.skills = skills
	c.inventoryCapacity = 10
	c.inventoryUpgrades = 0

}

func (c *Character) DisplayInfo() {
	fmt.Println("	INFOS			")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Name :", c.name)
	fmt.Println("Class :", c.class)
	fmt.Println("Level :", c.level)
	fmt.Println("Max Health :", c.health_max)
	fmt.Println("Current Health :", c.current_health)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
}

func (c Character) AccessInventory() {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("	INVENTORY			")
	fmt.Println("")
	fmt.Println("Max : ", c.inventoryCapacity)
	for i := range c.inventory {
		fmt.Println("You have", c.inventory[i], i, ".")
	}
	var choice int

	fmt.Println("1. Equip head")
	fmt.Println("2. Equip body")
	fmt.Println("3. Equip foot")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("9. Back to menu")

	fmt.Scan(&choice)

	switch choice {
	case 1:
	///Equipe head
	case 2:
		///Equip body
	case 3:
		///Equip boots

	}

}
