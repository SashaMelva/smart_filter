package entity

type Fileters struct {
	Filters []*Fileter
}

type Fileter struct {
	Name        string
	Description string
}
