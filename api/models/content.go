package models

type Category struct {
	GUID        string `json:"guid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	LangID      string `json:"lang_id"`
}

type CategoryList struct {
	Categories []Category `json:"categories"`
}

type Chapter struct {
	GUID        string `json:"guid"`
	CategoryID  string `json:"category_id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Description string `json:"description"`
	LangID      string `json:"lang_id"`
}

type ChapterList struct {
	Chapters []Chapter `json:"chapters"`
}

type Article struct {
	GUID  string `json:"guid"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Icon  string `json:"icon"`
	Media string `json:"media"`
}

type ArticleList struct {
	Articles []Article `json:"articles"`
}

type NewsListrequest struct {
	Status string `json:"status"`
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
}

type News struct {
	GUID        string `json:"guid"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	Image       string `json:"image"`
	PublishDate string `json:"publish_date"`
	Status      string `json:"status"`
	NextScreen  string `json:"next_screen"`
	LangID      string `json:"lang_id"`
}

type NewsList struct {
	News []News `json:"news"`
}

type ArticleA struct {
	ChapterID   string `json:"chapter_id"`
	LangID      string `json:"lang_id"`
	GUID        string `json:"guid"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	Icon        string `json:"icon"`
	Media       string `json:"media"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type Errors struct {
	StatusCode int64 `json:"status_code"`
	ResError   error `json:"res_error"`
}
