package red

import "fmt"

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