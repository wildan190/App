package app

type Base struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

func (b *Base) SetIDAndName(id int, name string) {
	b.ID = id
	b.Name = name
}
