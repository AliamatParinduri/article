package entity

type Model struct {
	Model interface{}
}

func RegisterModelPG() []Model {
	return []Model{
		{Model: User{}},
	}
}
