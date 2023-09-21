package red

import "fmt"

func(c *Character) Death() {
	if c.current_health <= 0 {
		fmt.Println("Wasted !")
		c.current_health = c.health_max / 2
		fmt.Println("You have been resurrected with 50% HP")
	} else {
		fmt.Println("Health Point :", c.current_health,"/", c.health_max)
	}
}