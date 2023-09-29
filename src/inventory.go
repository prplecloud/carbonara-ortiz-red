package red

import "fmt"

func (c *Character) DisplayInfo() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("	INFOS			")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Name :", c.name)
	fmt.Println("Class :", c.class)
	fmt.Println("Level :", c.level)
	fmt.Println("Exp :", c.currentExperience, "/", c.maxExperience)
	fmt.Println("Mana :", c.mana, "/", c.mana_max)
	fmt.Println("HP :", c.current_health, "/", c.health_max)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
}

func (c *Character) AccessInventory() {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("	INVENTORY			")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("")
	fmt.Println("Golds : ", c.golds)
	fmt.Println("Max : ", c.inventoryCapacity)

	for i := range c.inventory {
		fmt.Println("You have", c.inventory[i], i, ".")
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("------------------------------------")
	fmt.Println("1. Use a HP Potion")
	fmt.Println("2. Use a Potion poison")
	fmt.Println("3. Use a Mana Potion")
	fmt.Println("4. Use SpellBook : FireBall")
	fmt.Println("5. Equip Adventurer's Hat")
	fmt.Println("6. Equip Adventurer's Tunics")
	fmt.Println("7. Equip Adventurer's Boots")
	fmt.Println("8. Back to menu")
	fmt.Println("------------------------------------")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		c1.Takepot()
		fmt.Println("You used a HP Potion")
	case 2:
		c1.PoisonPot()
		fmt.Println("You used a Poison Potion")
	case 3:
		c1.TakepotMana()
		fmt.Println("You used a Mana Potion")
	case 4:
		c1.SpellBook("FireBall")
	case 5:
		c.EquipItem("Adventurer's hat")

	case 6:
		c.EquipItem("Adventurer's tunic")

	case 7:
		c.EquipItem("Adventurer's boots")
	}
}

func (m Shop) DisplayShop(c *Character) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Hello dear traveler.")
	fmt.Println("Would you like to buy an item ?")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Golds : ", c.golds)
	fmt.Println("1. Buy Hp Potion [Price : 3 golds]")
	fmt.Println("2. Buy Poison Potion [Price : 6 golds]")
	fmt.Println("3. Buy Mana Potion [Price : 5 golds]")
	fmt.Println("4. Buy Spellbook FireBall [Price : 25 golds]")
	fmt.Println("5. Buy Wolf skin [Price 4 golds]")
	fmt.Println("6. Buy Troll skin [Price 7 golds]")
	fmt.Println("7. Buy Wild boar leather [Price 3 golds]")
	fmt.Println("8. Buy Raven feather [Price 1 gold]")
	fmt.Println("9. Upgrade inventory slot by 10 [Price 30 golds]")
	fmt.Println("10. Leave")
	fmt.Println("------------------------------------")

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
		if c.golds >= 5 {
			c.golds -= 5
			c.AddInventory("Mana Potion")
		} else {
			fmt.Println("You don't have enough money")
		}
	case 4:
		if c.golds >= 25 {
			c.golds -= 25
			c.AddInventory("SpellBook : FireBall")
		} else {
			fmt.Println("You don't have enough money")
		}
	case 5:
		if c.golds >= 4 {
			c.golds -= 4
			c.AddInventory("Wolf skin")
		} else {
			fmt.Println("You don't have enough money")
		}
	case 6:
		if c.golds >= 7 {
			c.golds -= 7
			c.AddInventory("Troll skin")
		} else {
			fmt.Println("You don't have enough money")
		}
	case 7:
		if c.golds >= 3 {
			c.golds -= 3
			c.AddInventory("Wild boar leather")
		} else {
			fmt.Println("You don't have enough money")
		}

	case 8:
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
	default:
		fmt.Println("See you soon !")
	}
}

func (c Character) AddInventory(item string) {
	if c.MaxItem() {
		c.inventory[item] = c.inventory[item] + 1
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("You bought a", item)
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("")
		fmt.Println("")
		for i := range c.inventory {
			fmt.Println("You have", c.inventory[i], i, ".")
		}
	}
}

func (c Character) RemoveInventory(item string) {
	c.inventory[item] = c.inventory[item] - 1
	for i := range c.inventory {
		fmt.Println("You have", c.inventory[i], i, ".")
		if c.inventory[i] == 0 {
			delete(c.inventory, i)
		}
	}
}

func (c *Character) UpgradeInventorySlot() {
	if c.inventoryUpgrades < 3 {
		c.inventoryCapacity += 10
		c.inventoryUpgrades++
		fmt.Println("Inventory capacity increased to", c.inventoryCapacity)
	} else {
		fmt.Println("You can only upgrade your inventory capacity three times.")
	}
}

func (c *Character) MaxItem() bool {
	nbItem := 0
	for _, item := range c.inventory {
		nbItem += item
	}
	if nbItem >= c.inventoryCapacity {
		fmt.Println("Inventory is full, you cannot add more items.")
		return false

	} else {
		return true
	}
}
