package models

type CharacterInter interface {
	Restart()
	DrawCard()
	PlayCard()
}

var _ CharacterInter = (*Character)(nil)
