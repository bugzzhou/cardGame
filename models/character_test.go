package models

import (
	"fmt"
	"testing"
)

const DefaultDrawNum = 5

func InitCharacter() Character {
	Card1 := Card{
		ID: 1, Name: "攻击", CardType: "attack", Cost: 1,
	}
	Card2 := Card{
		ID: 2, Name: "防御", CardType: "defence", Cost: 1,
	}
	return Character{
		Name:    "a",
		Energy:  3,
		DrawNum: DefaultDrawNum,
		AllCards: []Card{
			Card1, Card1, Card1, Card1, Card1,
			Card2, Card2, Card2, Card2, Card2,
		},
		DrawCards: []Card{
			Card1, Card1, Card1, Card1, Card1,
			Card2, Card2, Card2, Card2, Card2,
		},
	}
}

func TestDrawCards(t *testing.T) {
	initC := InitCharacter()
	displayNum(initC)

	initC.DrawCard(5)
	displayNum(initC)

	initC.EndTurn()
	displayNum(initC)

}

func displayNum(c Character) {
	fmt.Printf("AllCards' num is: %d\n", len(c.AllCards))
	fmt.Printf("DrawCards' num is: %d\n", len(c.DrawCards))
	fmt.Printf("HandCards' num is: %d\n", len(c.HandCards))
	fmt.Printf("DiscardCards' num is: %d\n", len(c.DiscardCards))
}
