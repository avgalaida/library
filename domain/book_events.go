package domain

type Message interface {
	Key() string
}

const (
	KindBookCreated = iota + 1
	KindBookDeleted = iota + 1
)

type CreateBookDelta struct {
	//Kind        uint32 `json:"kind"`
	ID          string `json:"id"`
	Meta        string `json:"meta"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	Authors     string `json:"authors"`
	Description string `json:"desc"`
	CreatedAt   string `json:"createdAt"`
}

type DeleteBookDelta struct {
	//Kind   uint32 `json:"kind"`
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

func NewBookCreatedMessage(id, meta, status, title, authors, description, createdAt string) *CreateBookDelta {
	return &CreateBookDelta{
		//Kind:        KindBookCreated,
		ID:          id,
		Meta:        meta,
		Status:      status,
		Title:       title,
		Authors:     authors,
		Description: description,
		CreatedAt:   createdAt,
	}
}

func NewBookDeletedMessage(id, meta, status string) *DeleteBookDelta {
	return &DeleteBookDelta{
		//Kind:   KindBookDeleted,
		ID:     id,
		Meta:   meta,
		Status: status,
	}
}
