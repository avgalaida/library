package domain

type Message interface {
	Key() string
}

type CreateBookDelta struct {
	Type      string `json:"type"`
	ID        string `json:"id"`
	Meta      int    `json:"meta"`
	Status    string `json:"status"`
	Title     string `json:"title"`
	Authors   string `json:"authors"`
	CreatedAt string `json:"createdAt"`
}

type DeleteBookDelta struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Meta int    `json:"meta"`
}

type RestoreBookDelta struct {
	Type   string `json:"type"`
	ID     string `json:"id"`
	Meta   int    `json:"meta"`
	Status string `json:"status"`
}

func (m *CreateBookDelta) Key() string {
	return "книга.создана"
}

func (m *DeleteBookDelta) Key() string {
	return "книга.удалена"
}

func (m *RestoreBookDelta) Key() string {
	return "книга.восстановлена"
}
