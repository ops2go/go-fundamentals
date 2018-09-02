package structs

type Category struct {
	ID      int    `json:"category_id"`
	Name    string `json:"category_name"`
	Created string `json:"created_date"`
}

type Categories []Category

type CategoryCount struct {
	Name  string
	Count int
}
