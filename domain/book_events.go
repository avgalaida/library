package domain

type Message interface {
	Key() string
}

type CreateBookDelta struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	Meta        string `json:"meta"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	Authors     string `json:"authors"`
	Description string `json:"desc"`
	CreatedAt   string `json:"createdAt"`
}

type DeleteBookDelta struct {
	Type   string `json:"type"`
	ID     string `json:"id"`
	Meta   string `json:"meta"`
	Status string `json:"status"`
}

func (m *CreateBookDelta) Key() string {
	return "книга.создана"
}

func (m *DeleteBookDelta) Key() string {
	return "книга.удалена"
}
