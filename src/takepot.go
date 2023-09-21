package red

import "fmt"

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