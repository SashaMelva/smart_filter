package entity

type Fileters struct {
	Filters []*Fileter
}

type Fileter struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type FilterAgeCategores struct {
	Filters []*FilterAgeCategore
}

type FilterAgeCategore struct {
	Id          int    `json:"id"`
	Category_id string `json:"category_id"`
	Name        string `json:"name"`
	Params_eng  string `json:"params_eng"`
	Params_ru   string `json:"params_ru"`
}

type ProcentByCategoryUser struct {
	CategoryName string `json:"category_name"`
	Procent      int    `json:"procent"`
}

type ProcentByCategoresUser struct {
	Category []*ProcentByCategoryUser `json:"procent_category"`
}
