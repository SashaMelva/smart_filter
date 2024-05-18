package entity

type VideoId struct {
	Id            int      `json:"id"`
	Url           string   `json:"url"`
	Name          string   `json:"name"`
	AgeCategoryId int      `json:"age_category_id"`
	StatusId      int      `json:"status_id"`
	Tags          []string `json:"tags"`
}
type VideoIdStatus struct {
	Id     int `json:"id"`
	Status int `json:"status_id"`
}
type Video struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type VideoCheker struct {
	IdAccount int    `json:"id_account"`
	UrlVideo  string `json:"url"`
}

type AgeCategores struct {
	AgeCategores []*AgeCategory `json:"age_categores"`
}
type AgeCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type StatusVideos struct {
	AgeCategores []*StatusVideo `json:"status_videos"`
}

type StatusVideo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
