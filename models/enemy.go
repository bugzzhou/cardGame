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

	HandCards []Action
}

type Action struct {
}
