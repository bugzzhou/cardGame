package models

import (
	"fmt"
	"testing"
)

func InitCharacter() Character {
	Card1 := Card{
		ID: 1, Name: "攻击", CardType: "attack", Cost: 1,
	}
	Card2 := Card{
		ID: 2, Name: "防御", CardType: "defence", Cost: 1,
	}
	return Character{
		Name:   "a",
		Energy: 3,
		AllCards: []Card{
			Card1, Card1, Card1, Card1, Card1,
			Card2, Card2, Card2, Card2, Card2,
		},
	}
}

func TestDrawCards(t *testing.T) {
	initC := InitCharacter()
	fmt.Printf("initC is: %v\n", initC)
}


