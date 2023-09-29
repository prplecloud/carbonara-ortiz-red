package red

import "fmt"

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
func (c *Character) PlayerRound(g *Goblin) {

	fmt.Println("------------------------------------")
	fmt.Println(c.name)
	fmt.Println("HP :", c.current_health, "/", c.health_max)
	fmt.Println("Mana :", c.mana, "/", c.mana_max)
	fmt.Println("1. Normal attack  [-5 HP]")
	fmt.Println("2. SPELLS")
	fmt.Println("3. INVENTORY")
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

func (c *Character) HasEnoughMana(manaCost int) bool {
	return c.mana >= manaCost
}

func TrainingFight(c *Character, g *Goblin) {
	g.InitGoblin("Goblin", 40, 20, 5, 6)
	fmt.Printf("Player (%s) vs Monster (%s)\n", c.name, g.name)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	turn := 1
	if c.initiative > g.initiative {
		fmt.Print("\033[H\033[2J")
		for !(EnnemyDead(g)) && !(CharacDead(c)) {
			c.PlayerRound(g)
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
				c.PlayerRound(g)
				turn++
			}
		}
	}

	if EnnemyDead(g) {
		g.GobDeath(c)
	} else {
		c.Death()
	}
	Menu1(&c1, &m1, &g1)
}
