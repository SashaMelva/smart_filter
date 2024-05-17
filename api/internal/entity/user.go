package entity

type User struct {
	Id           int
	PhoneNumber  string
	FirstName    string
	MiddelName   string
	LastName     string
	Age          int
	DateBirthday int
	AgeCategory  int
}

type ChilgrenLists struct {
	IdUser    int
	Chilgrens []*Chilgren
}

type Chilgren struct {
	Id       int
	IdUser   int
	IdParent int
}

type FilterLists struct {
	IdUser    int
	Chilgrens []*Chilgren
}
