package entity

type User struct {
	Id           int    `json:"id"`
	AccountId    int    `json:"account_id"`
	PhoneNumber  string `json:"phone_number"`
	FirstName    string `json:"first_name"`
	MiddelName   string `json:"middel_name"`
	LastName     string `json:"last_name"`
	Age          int    `json:"age"`
	DateBirthday string `json:"date_birth"`
	AgeCategory  int    `json:"age_category"`
}

type UserCreater struct {
	PhoneNumber  string `json:"phone_number"`
	FirstName    string `json:"first_name"`
	MiddelName   string `json:"middel_name"`
	LastName     string `json:"last_name"`
	DateBirthday string `json:"date_birth"`
}

type ChilgrenLists struct {
	IdUser    int         `json:"parent_id"`
	Chilgrens []*Chilgren `json:"childrens"`
}

type Ids struct {
	Ids []int `json:"chaild_id"`
}

type Chilgren struct {
	IdUser     int    `json:"user_id"`
	FirstName  string `json:"first_name"`
	MiddelName string `json:"middel_name"`
	LastName   string `json:"last_name"`
}

type FilterLists struct {
	IdUser    int
	Chilgrens []*Chilgren
}
