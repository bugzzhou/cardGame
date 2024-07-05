package models

type CharacterInter interface {
	Restart()
	DrawCard(int)
	PlayCard(int, *Character)
	EndTurn()
	GetCard()
}

var _ CharacterInter = (*Character)(nil)
