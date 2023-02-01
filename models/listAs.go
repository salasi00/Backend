package models

type ListAs struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (ListAs) TableName() string {
	return "list_as"
}
