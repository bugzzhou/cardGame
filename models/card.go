package models

// TODO bugzzhou 将所有卡牌的效果做成一个函数;
// 现在写死
type Card struct {
	ID          int
	Name        string
	CardType    string
	Cost        int
	Description string
}

// TODO bugzzhou
type Lotion struct {
}

type Prop struct {
}

type Booty struct {
}
