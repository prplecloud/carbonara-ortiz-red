package red

import "fmt"

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