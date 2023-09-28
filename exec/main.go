package main

import (
	"fmt"
	"os"
	"time"
)

var choice int
var nickname string
var c1 Character
var m1 Shop
var g1 Goblin

type Character struct {
	name              string
	class             string
	level             int
	health_max        int
	current_health    int
	inventory         map[string]int
	skills            []string
	golds             int
	Equipment         string
	inventoryCapacity int
	inventoryUpgrades int
	initiative        int
	currentExperience int
	maxExperience     int
	mana              int
	mana_max          int
}

type Goblin struct {
	name           string
	health_max     int
	current_health int
	attack_point   int
	initiative     int
}

func (g *Goblin) InitGoblin(name string, health_max int, current_health int, attack_point int, initiative int) {
	g.name = name
	g.health_max = health_max
	g.current_health = current_health
	g.attack_point = attack_point
	g.initiative = initiative
}

type Equipment struct {
	head string
	body string
	feet string
}

func main() {

	m1.items = map[string]int{"Hp Potion": 1, "Poison Potion": 1, "SpellBook : FireBall": 1}
	m1.price = 0
	m1.quantity = 1
	fmt.Print("\033[H\033[2J")
	c1.SetName()
	fmt.Print("\033[H\033[2J")
	c1.ClassChoice()
	fmt.Print("\033[H\033[2J")
	for {
		fmt.Println("-------------------------------------")
		fmt.Println("Select an option")
		menu1(&c1, &m1, &g1)
	}
}

func (c *Character) Init(name string, class string, level int, health_max int, current_health int, inventory map[string]int, skills []string, golds int, initiative int, currentExperience int, maxExperience int, mana int, mana_max int) {
	c.name = name
	c.class = class
	c.level = level
	c.health_max = health_max
	c.current_health = current_health
	c.inventory = inventory
	c.skills = skills
	c.golds = golds
	c.initiative = initiative
	c.currentExperience = currentExperience
	c.maxExperience = maxExperience
	c.mana = mana
	c.mana_max = mana_max

	c.inventoryCapacity = 10
	c.inventoryUpgrades = 0
}

func menu1(c1 *Character, m1 *Shop, g1 *Goblin) {
	fmt.Println("-------------------------------------")
	fmt.Println("|	     MAIN MENU 		    |")
	fmt.Println("|1. Display Character Informations  |")
	fmt.Println("|2. Display Inventory		    |")
	fmt.Println("|3. Display Shop		    |")
	fmt.Println("|4. Black-Smith	    		    |")
	fmt.Println("|5. Show equiped Equipment	    |")
	fmt.Println("|6. Training Fight		    |")
	fmt.Println("|15. Quit Game			    |")
	fmt.Println("-------------------------------------")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		c1.DisplayInfo()
	case 2:
		c1.AccessInventory()
	case 3:
		m1.DisplayShop(c1)
	case 4:
		c1.BlackSmith()
	case 5:

	case 6:
		TrainingFight(c1, g1)
	case 7:

	case 15:
		os.Exit(0)
		///Quit
	default:
		fmt.Println("------------------------------------")
		fmt.Println("Choice is not valid. Please choose a valid option")
	}
}

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

func (c Character) AccessInventory() {
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
	fmt.Println("5. Back to menu")
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
		if c.current_health+healthpoint <= c.health_max {
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
func (c *Character) TakepotMana() {
	const manapoint int = 20
	if c.inventory["Mana Potion"] == 0 {
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("You have no Mana Potion")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("")
		fmt.Println("")

		return
	}
	if c.mana == c.mana_max {
		fmt.Println("You already are full mana")
		return
	}
	if c.inventory["Mana Potion"] > 0 {
		c.inventory["Mana Potion"]--
		if c.mana+manapoint <= c.mana_max {
			c.mana += manapoint
		} else {
			c.mana = c.mana_max
			return
		}
		fmt.Println("Mana set to :", c.mana, "/", c.mana_max)
		if c.inventory["Mana Potion"] == 0 {
			delete(c.inventory, "Mana Potion")
		}
	}
}

func (c *Character) Death() {
	if c.current_health <= 0 {
		fmt.Println("Wasted !")
		c.current_health = c.health_max / 2
		fmt.Println("You have been resurrected with 50% HP")
		fmt.Println("You lost the fight")
	} else {
		fmt.Println("Health Point :", c.current_health, "/", c.health_max)
	}
}
func (g *Goblin) GobDeath(c *Character) {
	if g.current_health <= 0 {
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("Goblin died")
		fmt.Println("You won the fight")
		experienceGain := c.Experience(g)
		c1.currentExperience += experienceGain
		fmt.Printf("You earned %d exp.\n", experienceGain)
		if c.currentExperience >= c.maxExperience {
			c.LevelUp()
		}
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		menu1(&c1, &m1, &g1)
	}
}

type Shop struct {
	items    map[string]int
	price    int
	quantity int
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

func (c *Character) PoisonPot() {
	if c.inventory["Poison Potion"] == 0 {
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("You have no Poison Potion")
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
		fmt.Println("")
		fmt.Println("")
		return
	}
	c.skills = append(c.skills, skills)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Skill added to your Spellbook")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("")
	fmt.Println("")
}

func (c *Character) SetName() {
	fmt.Println("Enter your characters name")
	fmt.Scanln(&nickname)
	if IsAlpha(nickname) {
		nickname = Capitalize(nickname)
		fmt.Println("Your characters name is : ", nickname)
	} else {
		fmt.Println("Your characters name must only contain letters")
		c.SetName()
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

	fmt.Println("What class do you want to choose ?")
	fmt.Println("1. Human")
	fmt.Println("2. Elf")
	fmt.Println("3. Dwarf")

	fmt.Scan(&choice)
	switch choice {
	case 1:
		c.Init(nickname, "Human", 1, 100, 50, map[string]int{"Sword": 1, "Hp Potion": 3, "Dark Hood": 1}, []string{"Coup de poing"}, 100, 10, 0, 100, 20, 40)
	case 2:
		c.Init(nickname, "Elf", 1, 80, 40, map[string]int{"Dagger": 1, "Hp Potion": 3, "Dark Hood": 1}, []string{"Coup de poing"}, 100, 10, 0, 100, 20, 40)
	case 3:
		c.Init(nickname, "Dwarf", 1, 120, 60, map[string]int{"Dagger": 1, "Hp Potion": 3, "Dark Hood": 1}, []string{"Coup de poing"}, 100, 10, 0, 100, 20, 40)
	default:
		fmt.Println("------------------------------------")
		fmt.Println("Choice is not valid. Please choose a valid option")
		c.ClassChoice()
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

func (c *Character) playerRound(g *Goblin) {

	fmt.Println("------------------------------------")
	fmt.Println(c.name)
	fmt.Println("HP :", c.current_health, "/", c.health_max)
	fmt.Println("Mana :", c.mana, "/", c.mana_max)
	fmt.Println("1. Normal attack  [-5 HP]")
	fmt.Println("2. SPELLS")
	fmt.Println("3. INVENTORY")
	fmt.Println("4. Run Away")
	fmt.Println("------------------------------------")

	fmt.Scan(&choice)
	switch choice {
	case 1:
		g.current_health -= 5
	case 2:
		c.Attack(g)
	case 3:
		c.AccessInventory()
	default:
		fmt.Println("------------------------------------")
		fmt.Println("Choice is not valid. Please choose a valid option")
	}
}

func (g *Goblin) GoblinRound(c *Character, turn int) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Turn :", turn)
	fmt.Println(g.name, ":", g.current_health, "/", g.health_max, "HP")
	fmt.Println(c.name, ":", c.current_health, "/", c.health_max, "HP")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("")
	fmt.Println("")
	if turn%3 == 0 {
		attack_point := g1.attack_point * 2
		fmt.Println("")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("The Monster attacks -10 HP")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("")
		c.current_health -= attack_point
		fmt.Println(g.name, ":", g.current_health, "/", g.health_max, "HP")
	} else {
		fmt.Println("")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("The Monster attacks -5 HP")
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("")
		fmt.Println("")
		c.current_health -= g.attack_point
	}

}

func EnnemyDead(g *Goblin) bool {
	if g.current_health <= 0 {
		return true
	}
	return false
}

func CharacDead(c *Character) bool {
	if c.current_health <= 0 {
		return true
	}
	return false
}

func TrainingFight(c *Character, g *Goblin) {
	g.InitGoblin("Goblin", 40, 20, 5, 6)
	fmt.Printf("Joueur (%s) vs Monstre (%s)\n", c.name, g.name)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	turn := 1
	if c.initiative > g.initiative {
		fmt.Print("\033[H\033[2J")
		for !(EnnemyDead(g)) && !(CharacDead(c)) {
			c.playerRound(g)
			if !(EnnemyDead(g)) || !(CharacDead(c)) {
				g.GoblinRound(c, turn)
				turn++
			}
		}
	} else {
		fmt.Print("\033[H\033[2J")
		for !(EnnemyDead(g)) && !(CharacDead(c)) {
			g.GoblinRound(c, turn)
			if !(EnnemyDead(g)) || !(CharacDead(c)) {
				c.playerRound(g)
				turn++
			}
		}
	}

	if EnnemyDead(g) {
		g.GobDeath(c)
	} else {
		c.Death()
	}

	menu1(&c1, &m1, &g1)
}

func (c *Character) Attack(g *Goblin) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("------------------------------------")
	fmt.Println("Choose an attack")
	fmt.Println("1.Direct hit [-8 HP, -3 mana]")
	fmt.Println("2.FireBall [-18 HP, -15 mana]")
	fmt.Println("------------------------------------")
	fmt.Println("")
	fmt.Println("")

	fmt.Scan(&choice)
	switch choice {
	case 1:
		if c.HasEnoughMana(3) {
			fmt.Println("You use Direct hit")
			fmt.Println("You used 3 mana")
			g.current_health -= 8
			fmt.Println(g.name, "took 8 damages")
			c.mana -= 0
		} else {
			fmt.Println("Not enough mana to cast this spell.")
		}
	case 2:
		if c.HasEnoughMana(15) {
			fmt.Println("You used FireBall")
			fmt.Println("You used 15 mana")
			g.current_health -= 18
			fmt.Println(g.name, "took 18 damages")
			c.mana -= 15
		} else {
			fmt.Println("Not enough mana to cast this spell.")
		}
	}
}

func (c *Character) Experience(g *Goblin) int {

	experienceGain := g.health_max / 2
	return experienceGain
}

func (c *Character) LevelUp() {
	c.level++
	c.maxExperience *= 2

	c.currentExperience = 0

	fmt.Printf("Level up ! You  now are level %d !\n", c.level)
}

func (c *Character) HasEnoughMana(manaCost int) bool {
	return c.mana >= manaCost
}
