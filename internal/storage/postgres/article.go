package postgres

type Article struct {
	Model
	Icon      string  `json:"icon"`
	Title     string  `json:"title"`
	Color     string  `json:"color"`
	IsDefault bool    `json:"is_default"`
	UserID    *uint64 `json:"user_id"`
	User      *User   `json:"user"`
}

func createDefaultArticles() {
	articles := []Article{
		{Icon: "home", Title: "Дом", Color: "green.500", IsDefault: true},
		{Icon: "cafe", Title: "Кафе", Color: "red.500", IsDefault: true},
		{Icon: "medicine", Title: "Здоровье", Color: "blue.500", IsDefault: true},
		{Icon: "transport", Title: "Транспорт", Color: "purple.500", IsDefault: true},
		{Icon: "family", Title: "Семья", Color: "yellow.500", IsDefault: true},
		{Icon: "student", Title: "Образование", Color: "cyan.500", IsDefault: true},
		{Icon: "basket", Title: "Продукты", Color: "pink.500", IsDefault: true},
		{Icon: "gift", Title: "Подарок", Color: "teal.500", IsDefault: true},
		{Icon: "sport", Title: "Спорт", Color: "orange.500", IsDefault: true},
		{Icon: "money", Title: "Зарплата", Color: "teal.700", IsDefault: true},
	}

	var count int64

	query := Db.Table("articles")
	query = query.Where("is_default = ?", true)
	query.Count(&count)

	if count == 0 {
		Db.Create(articles)
	}

}
