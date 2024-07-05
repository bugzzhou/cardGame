package models

// TODO bugzzhou 将所有卡牌的效果做成一个函数;
// 现在写死
type Card struct {
	ID          int
	Level       int //表示是否升级过了
	Name        string
	CardType    string //TODO bugzzhou 可能需要增加一个字段，用于表示虚无、消耗等
	Cost        int
	AffectFunc  AffectFuncInter //会将作战双方的全量传递进去，读取对应buff+debuff；再根据不同的FuncId，实现不同的函数
	Description string
}

// TODO bugzzhou

type Effect struct {
	TypeId    int //状态类型-唯一标识
	Value     int //表示数值-专指buff；debuff该字段不生效
	LastTurns int //持续回合（永久则用-1表示-每回合需要减一）
}

type Lotion struct {
}

type Prop struct {
}

type Booty struct {
}
