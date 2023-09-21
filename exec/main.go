package main

import ("fmt"
		"time"
		"os"
)

type Character struct {
	name string
	class string
	level int
	health_max int
	current_health int
	inventory map[string]int
	skills []string
	golds int
	Equipment string
	inventoryCapacity int
	inventoryUpgrades int
}

type Equipment struct {
	head string
	body string
	feet string
}

var nickname string

func main() {		
	var c1 Character
	var m1 Shop

	m1.items = map[string]int{"Hp Potion" : 1, "Poison Potion" : 1, "SpellBook : FireBall" : 1}
	m1.price = 0
	m1.quantity = 1
	
	c1.SetName()
	c1.ClassChoice()
	for {
		fmt.Println("-------------------------------------")
		fmt.Println("Character : ", nickname)
		fmt.Println("Class : ", c1.class)
		fmt.Println("Golds : ", c1.golds)
		fmt.Println("Inventory Capacity :", c1.inventoryCapacity)
		c1.Death()
		fmt.Println("Select an option")
		menu1(&c1, &m1)			
	}
}






func (c *Character) Init(name string, class string, level int, health_max int, current_health int, inventory map[string]int, skills []string, golds int) {
	c.name = name
	c.class = class
	c.level = level
	c.health_max = health_max
	c.current_health = current_health
	c.inventory = inventory
	c.skills = skills
	c.golds = golds

	c.inventoryCapacity = 10
    c.inventoryUpgrades = 0
}





func menu1(c1 *Character, m1 *Shop) {

	var choice int

fmt.Println("-------------------------------------")
fmt.Println("|	     MAIN MENU 		    |")
fmt.Println("|1. Display Character Informations  |")
fmt.Println("|2. Display Inventory		    |")
fmt.Println("|3. Verify HP		   	    |")
fmt.Println("|4. Take a Hp Potion	  	    |")
fmt.Println("|5. Take a Poison Potion	    |")
fmt.Println("|6. Display Shop		    |")
fmt.Println("|7. Use SpellBook : FireBall	    |")
fmt.Println("|8. Black-Smith	    		    |")
fmt.Println("|9. Show equiped Equipment	    		    |")
fmt.Println("|15. Quit Game			    |")
fmt.Println("-------------------------------------")
fmt.Println("")
fmt.Println("")
fmt.Println("")

fmt.Scan(&choice)

	switch choice {
	case 1:
		c1.DisplayInfo()
		///DisplaycharacInfo
	case 2:
		c1.AccessInventory()
		///DisplayInventory
	case 3:
		c1.Death()
		///Vérifier les HP 
	case 4:
		c1.Takepot()
		///Prendre une potion de HP
	case 5:
		c1.PoisonPot()
	case 6:
		m1.DisplayShop(c1)
	case 7:
		c1.SpellBook("FireBall")
		///Utilise SpellBook
	case 8:
		c1.BlackSmith()
		///Montre Forgeron
	case 9:
		///Show equipment
		
	case 15 :
		os.Exit(0)
		///Quit
	default :
	fmt.Println("------------------------------------")
		fmt.Println("Choice is not valid. Please choose a valid option")
	}
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




func (c *Character) Takepot() {
	const healthpoint int = 50
	if c.inventory["Hp Potion"] == 0 {
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("You have no HP Potion")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("")
		fmt.Println("")
		return
	}
	if c.current_health == c.health_max {
		fmt.Println("You already are full life")
		return
	}
	if c.inventory["Hp Potion"] > 0 {
		 c.inventory["Hp Potion"]--
		if c.current_health + healthpoint <= c.health_max {
	 	c.current_health += healthpoint
		} else {
			c.current_health = c.health_max
			return
		}
		fmt.Println("Hp set to :", c.current_health, "/", c.health_max)
		if c.inventory["Hp Potion"] == 0 {
			delete(c.inventory, "Hp Potion")
		}
	}
}




func(c *Character) Death() {
	if c.current_health <= 0 {
		fmt.Println("Wasted !")
		c.current_health = c.health_max / 2
		fmt.Println("You have been resurrected with 50% HP")
	} else {
		fmt.Println("Health Point :", c.current_health,"/", c.health_max)
	}
}

type Shop struct {

	items map[string]int
	price int 
	quantity int
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
	if c.MaxItem() {
		c.inventory[item] = c.inventory[item] + 1
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("You bought a", item,)
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
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




func(c *Character) PoisonPot() {
	if c.inventory["Poison Potion"] == 0 {
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("You have no Poison poison")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("")
		fmt.Println("")
	}
	if c.inventory["Poison Potion"] > 0 {
		c.inventory["Poison Potion"]--
		fmt.Println("You used a Poison Potion")
		for i := 0; i < 3; i++ {
			c.current_health -= 10
			time.Sleep(1 * time.Second)
			fmt.Println("------------------------------------")
			fmt.Println(c.current_health, "/", c.health_max, "Health Point")
			if c.inventory["Poison Potion"] == 0 {
				delete(c.inventory, "Poison Potion")
			}
		}
	}
}




func (c *Character) SpellBook(skills string) {
    var HasSkill bool
    for _, CharSkill := range c.skills {
        if skills == CharSkill {
            HasSkill = true
            break
        }
		if c.inventory["SpellBook : FireBall"] == 0 {
			fmt.Println("You have no SpellBook : FireBall")
			return
		}
		if c.inventory["SpellBook : FireBall"] > 0 {
			c.inventory["SpellBook : FireBall"]--
		}
    }
    if HasSkill {
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
        fmt.Println("You already have this skill")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
        return	
    }
    c.skills = append(c.skills, skills)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
    fmt.Println("Skill added to your Spellbook")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}






func (c *Character) SetName() {
	fmt.Println("Enter your characters name")
	fmt.Scanln(&nickname)
	if IsAlpha(nickname) {
		nickname = Capitalize(nickname)
		fmt.Println("Your characters name is : ", nickname)
	} else {
		fmt.Println("Your characters name must only contain letters")
	}
}





func IsAlpha(s string) bool {
	h := []rune(s)
	for i := 0; i <= len(h)-1; i++ {
		if (h[i] >= 0) && (h[i] <= 64) || (h[i] >= 91) && (h[i] <= 96) || (h[i] >= 123) && (h[i] <= 127) {
			return false
		}
	}
	return true
}





func prim(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}




func Capitalize(s string) string {
	ar := []rune(s)
	letra := true
	for i := 0; i < len(s); i++ {
		if prim(ar[i]) == true && letra {
			if ar[i] >= 'a' && ar[i] <= 'z' {
				ar[i] = 'A' - 'a' + ar[i]
			}
			letra = false
		} else if ar[i] >= 'A' && ar[i] <= 'Z' {
			ar[i] = 'a' - 'A' + ar[i]
		} else if prim(ar[i]) == false {
			letra = true
		}
	}
	return string(ar)
}




func (c *Character) ClassChoice() {
	var choice int
	
	fmt.Println("What class do you want to choose ?")
	fmt.Println("1. Human")
	fmt.Println("2. Elf")
	fmt.Println("3. Dwarf")

	fmt.Scan(&choice)
switch choice {
case 1:
	c.Init(nickname, "Human", 1, 100, 50, map[string]int {"Sword" : 3, "Hp Potion" : 1, "Dark Hood" : 1}, []string{"Coup de poing"}, 10000)
case 2:
	c.Init(nickname, "Elf", 1, 80, 40, map[string]int {"Dagger" : 1, "Hp Potion" : 1, "Dark Hood" : 1}, []string{"Coup de poing"}, 100)
case 3:
	c.Init(nickname, "Dwarf", 1, 120, 60, map[string]int {"Dagger" : 1, "Hp Potion" : 1, "Dark Hood" : 1}, []string{"Coup de poing"}, 100)
	default :
	fmt.Println("------------------------------------")
		fmt.Println("Choice is not valid. Please choose a valid option")
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



func (c *Character) UpgradeInventorySlot() {
	if c.inventoryUpgrades < 3 {
		c.inventoryCapacity += 10
		c.inventoryUpgrades++
		fmt.Println("Inventory capacity increased to", c.inventoryCapacity)
	} else {
		fmt.Println("You can only upgrade your inventory capacity three times.")
	}
}



func (c Character) BlackSmith() {

	var choice int

	fmt.Println("1. Craft an Adventurer's hat")
	fmt.Println("2. Craft an Adventurer's tunic")
	fmt.Println("3. Craft an Adventurer's boots")

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
			c.inventory["Wolf skin"]-= 2
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
		default :
	fmt.Println("------------------------------------")
		fmt.Println("Choice is not valid. Please choose a valid option")
	}
}







