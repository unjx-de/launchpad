package bookmark

type Bookmark struct {
	Name string `json:"name" validate:"required"`
	Icon string `json:"icon" validate:"required"`
	Url  string `json:"url" validate:"required"`
}
