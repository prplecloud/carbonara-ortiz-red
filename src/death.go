package red

import "fmt"

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
		Menu1(&c1, &m1, &g1)
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