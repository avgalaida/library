package domain

type BookView struct {
	ID        string `json:"id"`
	Meta      int    `json:"meta"`
	Status    string `json:"status"`
	Title     string `json:"title"`
	Authors   string `json:"authors"`
	CreatedAt string `json:"createdAt"`
}

func NewBookView(b Book) BookView {
	return BookView{
		ID:        b.Base.ID,
		Meta:      b.Base.Meta,
		Status:    b.Status,
		Title:     b.Title,
		Authors:   b.Authors,
		CreatedAt: b.Base.CreatedAt,
	}
}
