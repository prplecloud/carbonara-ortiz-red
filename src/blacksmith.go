package red

import "fmt"

func (c Character) BlackSmith() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("			BLACKSMITH")
	fmt.Println("1. Craft an Adventurer's hat [1 Raven feather, 1 wild boar leather]")
	fmt.Println("2. Craft an Adventurer's tunic [2 Wolf skin, 1 Troll skin]")
	fmt.Println("3. Craft an Adventurer's boots [1 Wolf skin, 1 wild boar leather]")
	fmt.Println("4. Leave")

	fmt.Scan(&choice)

	switch choice {

	case 1:
		if c.inventory["Raven feather"] > 0 && c.inventory["Wild boar leather"] > 0 {
			c.inventory["Raven feather"]--
			c.inventory["Wild boar leather"]--
			c.AddInventory("Adventurer's hat")
			c.golds -= 5
			fmt.Println("You crafted an Adventurer's hat")
		} else {
			fmt.Println("You don't have the required element in your inventory")
		}
		if c.inventory["Raven feather"] == 0 {
			delete(c.inventory, "Raven feather")
		}
	case 2:
		if c.inventory["Wolf skin"] > 1 && c.inventory["Troll skin"] > 0 {
			c.inventory["Wolf skin"] -= 2
			c.inventory["Troll skin"]--
			c.AddInventory("Adventurer's tunic")
			c.golds -= 5
			fmt.Println("You crafted an Adventurer's tunic")
		} else {
			fmt.Println("You don't have the required element in your inventory")
		}

	case 3:
		if c.inventory["Wolf skin"] > 0 && c.inventory["Wild boar leather"] > 0 {
			c.inventory["Wolf skin"]--
			c.inventory["Wild boar leather"]--
			c.AddInventory("Adventurer's boots")
			c.golds -= 5
			fmt.Println("You crafted an Adventurer's boots")
		} else {
			fmt.Println("You don't have the required element in your inventory")
		}
	default:
		fmt.Println("------------------------------------")
		fmt.Println("Choice is not valid. Please choose a valid option")
	}
}