package models

type CharacterInter interface {
	Restart()
	DrawCard(int)
	PlayCard(int)
	EndTurn()
	GetCard()
}

var _ CharacterInter = (*Character)(nil)
