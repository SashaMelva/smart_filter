package entity

type Account struct {
	Login    string
	Password string
	Role     int
}

type AccountId struct {
	Id       int
	Login    string
	Password string
	Role     int
}
