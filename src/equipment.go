package red

type Equipment struct {
	head string
	body string
	feet string
}

func (c *Character) Equipped() {

	if c.equipment.head != "" {
		c.health_max += 10
	}
	if c.equipment.body != "" {
		c.health_max += 25
	}
	if c.equipment.feet != "" {
		c.health_max += 15
	}
}
