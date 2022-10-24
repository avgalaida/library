package domain

type BookView struct {
	ID          string `json:"id"`
	Meta        string `json:"meta"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	Authors     string `json:"authors"`
	Description string `json:"desc"`
	CreatedAt   string `json:"createdAt"`
}

func NewBookView(b Book) BookView {
	return BookView{
		ID:          b.Base.ID,
		Meta:        b.Base.Meta,
		Status:      b.Status,
		Title:       b.Title,
		Authors:     b.Authors,
		Description: b.Description,
		CreatedAt:   b.Base.CreatedAt,
	}
}
