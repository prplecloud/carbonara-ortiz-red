package red

type Character struct {
	name              string
	class             string
	level             int
	health_max        int
	current_health    int
	inventory         map[string]int
	skills            []string
	golds             int
	equipment         Equipment
	inventoryCapacity int
	inventoryUpgrades int
	initiative        int
	currentExperience int
	maxExperience     int
	mana              int
	mana_max          int
}

type Goblin struct {
	name           string
	health_max     int
	current_health int
	attack_point   int
	initiative     int
}
type Equipment struct {
	head string
	body string
	feet string
	}

type Shop struct {
	items    map[string]int
	price    int
	quantity int
	}