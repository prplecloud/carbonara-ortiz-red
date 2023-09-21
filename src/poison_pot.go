package red

import ("time"
		"fmt"
)

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