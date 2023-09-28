package red

import "fmt"

var choice int
var nickname string
var c1 Character
var m1 Shop
var g1 Goblin


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

func (c *Character) Init(name string, class string, level int, health_max int, current_health int, inventory map[string]int, skills []string, golds int, initiative int, currentExperience int, maxExperience int, mana int, mana_max int, equipment Equipment) {
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
	c.equipment = equipment
	c.inventoryCapacity = 10
	c.inventoryUpgrades = 0
}

func (g *Goblin) InitGoblin(name string, health_max int, current_health int, attack_point int, initiative int) {
	g.name = name
	g.health_max = health_max
	g.current_health = current_health
	g.attack_point = attack_point
	g.initiative = initiative
}

func (c *Character) ClassChoice() {

	fmt.Println("What class do you want to choose ?")
	fmt.Println("1. Human")
	fmt.Println("2. Elf")
	fmt.Println("3. Dwarf")

	fmt.Scan(&choice)
	switch choice {
	case 1:
		c.Init(nickname, "Human", 1, 100, 50, map[string]int{"Sword": 1, "Hp Potion": 3, "Dark Hood": 1}, []string{"Coup de poing"}, 100, 10, 0, 100, 20, 40, c.equipment)
	case 2:
		c.Init(nickname, "Elf", 1, 80, 40, map[string]int{"Dagger": 1, "Hp Potion": 3, "Dark Hood": 1}, []string{"Coup de poing"}, 100, 10, 0, 100, 20, 40, c.equipment)
	case 3:
		c.Init(nickname, "Dwarf", 1, 120, 60, map[string]int{"Dagger": 1, "Hp Potion": 3, "Dark Hood": 1}, []string{"Coup de poing"}, 100, 10, 0, 100, 20, 40, c.equipment)
	default:
		fmt.Println("------------------------------------")
		fmt.Println("Choice is not valid. Please choose a valid option")
		c.ClassChoice()
	}
}

