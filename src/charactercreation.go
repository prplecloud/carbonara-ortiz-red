package red

import "fmt"

func SetName() {
	fmt.Println("Enter your character name")
	var name string		
	fmt.Scanln(&name)
		if IsAlpha(name) {
		fmt.Println("Your character name is", name)
	} else {
		fmt.Println("Your character name must only contain letters")
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

func (c *Character) ClassChoice() {

	var choice int
	
	fmt.Println("What class do you want to choose ?")
	fmt.Println("1. Human")
	fmt.Println("2. Elf")
	fmt.Println("3. Dwarf")

	fmt.Scan(&choice)
switch choice {
case 1:
	c.Init(nickname, "Human", 1, 100, 50, map[string]int {"Sword" : 1, "Hp Potion" : 1, "Dark Hood" : 1}, []string{"Coup de poing"}, 100)
case 2:
	c.Init(nickname, "Elf", 1, 80, 40, map[string]int {"Dagger" : 1, "BBBBB" : 1, "Dark Hood" : 1}, []string{"Coup de poing"}, 100)
case 3:
	c.Init(nickname, "Dwarf", 1, 120, 60, map[string]int {"Dagger" : 1, "Hp Potion" : 1, "CCCC" : 1}, []string{"Coup de poing"}, 100)
	default :
	fmt.Println("------------------------------------")
		fmt.Println("Choice is not valid. Please choose a valid option")
	}
}

