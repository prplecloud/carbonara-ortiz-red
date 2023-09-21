package red

import "fmt"

type Shop struct {

	items map[string]int
	price int 
	quantity int
}

func mainmerch() {
	var m1 Shop

	m1.items = map[string]int{"Hp Potion" : 1, "Poison Potion" : 1, "SpellBook : FireBall" : 1}
	m1.price = 0
	m1.quantity = 1
}

func (m Shop) DisplayShop(c *Character) {
	
	var choice int
	fmt.Println("------------------------------------")
	fmt.Println("Hello dear traveler.")
	fmt.Println("Would you like to buy an item ?")
	fmt.Println("------------------------------------")
	fmt.Println("1. Buy Hp Potion [Price : 3 golds]")
	fmt.Println("2. Buy Poison Potion [Price : 6 golds]")
	fmt.Println("3. Buy Spellbook FireBall [Price : 25 golds]")
	fmt.Println("4. Buy Wolf skin [Price 4 golds]")
	fmt.Println("5. Buy Troll skin [Price 7 golds]")
	fmt.Println("6. Buy Wild boar leather [Price 3 golds]")
	fmt.Println("7. Buy Raven feather [Price 1 gold]")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("15. Upgrade inventory slot by 10 [Price 30 golds]")
	fmt.Println("Anything else. Leave")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")


	fmt.Scan(&choice)
	switch choice {
	case 1:
		if c.golds >= 3 {
			c.golds -= 3
		c.AddInventory("Hp Potion")
		} else {
			fmt.Println("You don't have enough money")
		}
	case 2:
		if c.golds >= 6 {
			c.golds -= 6
			c.AddInventory("Poison Potion")
		} else {
			fmt.Println("You don't have enough money")
		}
	case 3:
		if c.golds >= 25 {
			c.golds -= 25
			c.AddInventory("SpellBook : FireBall")
		} else {
			fmt.Println("You don't have enough money")
		}
	case 4:
		if c.golds >= 4 {
			c.golds -= 4
			c.AddInventory("Wolf skin")
		} else {
			fmt.Println("You don't have enough money")
		}
	case 5:
		if c.golds >= 7 {
			c.golds -= 7
			c.AddInventory("Troll skin")
		} else {
			fmt.Println("You don't have enough money")
		}
	case 6:
		if c.golds >= 3 {
			c.golds -= 3
			c.AddInventory("Wild boar leather")
		} else {
			fmt.Println("You don't have enough money")
		}

	case 7:
		if c.golds >= 1 {
			c.golds -= 1
			c.AddInventory("Raven feather")
		} else {
			fmt.Println("You don't have enough money")
		}
	case 15:
		if c.golds >= 30 {
			c.golds -= 30
			c.UpgradeInventorySlot()
		
		} else {
			fmt.Println("You don't have enough money")
		}
		default : fmt.Println("See you soon !")
	}
}

func (c Character) AddInventory(item string) {
	c.inventory[item] = c.inventory[item] + 1
	fmt.Println("------------------------------------")
	fmt.Println("You bought", item)
	for i := range c.inventory { 
		fmt.Println("You have", c.inventory[i], i, ".")
	}

}

func (c Character) RemoveInventory(item string) {
	c.inventory[item] = c.inventory[item] - 1 
	fmt.Println("------------------------------------")
	fmt.Println("You sold", item)
		for i := range c.inventory { 
			fmt.Println("You have", c.inventory[i], i, ".")
			if c.inventory["Hp Potion"] == 0 {
			delete(c.inventory, "Hp Potion")
		}
	}
}
