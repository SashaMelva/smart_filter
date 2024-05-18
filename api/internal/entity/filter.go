package entity

type Fileters struct {
	Filters []*Fileter
}

type Fileter struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
