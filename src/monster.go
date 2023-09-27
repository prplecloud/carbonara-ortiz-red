package red

type Monster struct {
	name           string
	health_max     int
	current_health int
	attack_point   int
	turn           int
	initiative     int
}

func (t Monster) MainGoblin() {
	t.InitGoblin("Goblin dummy", 40, 40, 5)
}
func (t Monster) InitGoblin(name string, health_max int, current_health int, attack_point int) {
	t.name = name
	t.health_max = health_max
	t.current_health = current_health
	t.attack_point = attack_point

}
