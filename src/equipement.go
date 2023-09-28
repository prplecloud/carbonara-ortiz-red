package red

import "fmt"

func (c *Character) EquipItem(item string) {
	if _, ok := c.inventory[item]; !ok {
		fmt.Println("You don't have", item, "in your inventory.")
		return
	}
	switch item {
	case "Adventurer's hat":
		if c.equipment.head == "" {
			c.equipment.head = item
			c.health_max += 10
			fmt.Println("You equipped Adventurer's hat")
		} else {
			fmt.Println("You already have a headgear equipped.")
			return
		}
	case "Adventurer's tunic":
		if c.equipment.body == "" {
			c.equipment.body = item
			c.health_max += 25
			fmt.Println("You equipped Adventurer's tunic")
		} else {
			fmt.Println("You already have a body armor equipped.")
			return
		}
	case "Adventurer's boots":
		if c.equipment.feet == "" {
			c.equipment.feet = item
			c.health_max += 15
			fmt.Println("You equipped Adventurer's boots")
		} else {
			fmt.Println("You already have boots equipped.")
			return
		}
	default:
		fmt.Println("Item not recognized")
	}
}