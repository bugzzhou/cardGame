package models

type Character struct {
	Name         string
	Energy       int
	AllCards     []Card
	DrawCards    []Card
	DiscardCards []Card
	HandCards    []Card
	Lotions      []Lotion
	Props        []Prop
}

func (c Character) PlayCard() {}

func (c Character) DrawCard() {}

func (c Character) Restart() {}
