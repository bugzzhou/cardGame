package models

type Enemy struct {
	Name        string
	HealthLimit int
	Health      int

	Power     int
	Dexterity int
	Shield    int

	Buffs   []Effect
	Debuffs []Effect

	Action []Action
}

type Action struct {
}
