package red

import ("fmt"
	"time")

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
	if  c.inventory["Hp Potion"] > 0 {
		c.inventory["Hp Potion"]--
		if c.current_health + healthpoint <= c.health_max {
	 	c.current_health += healthpoint
		} else {
			c.current_health = c.health_max
			return
		}
		fmt.Println("Hp set to :", c.current_health, "/", c.health_max)
		if  c.inventory["Hp Potion"] == 0 {
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