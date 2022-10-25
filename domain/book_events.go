package domain

type Message interface {
	Key() string
}

type CreateBookDelta struct {
	Type      string `json:"type"`
	ID        string `json:"id"`
	Status    string `json:"status"`
	Title     string `json:"title"`
	Authors   string `json:"authors"`
	CreatedAt string `json:"createdAt"`
}
type DeleteBookDelta struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
type RestoreBookDelta struct {
	Type   string `json:"type"`
	ID     string `json:"id"`
	Status string `json:"status"`
}
type ChangeBookTitleDelta struct {
	Type  string `json:"type"`
	ID    string `json:"id"`
	Title string `json:"title"`
}
type ChangeBookAuthorsDelta struct {
	Type    string `json:"type"`
	ID      string `json:"id"`
	Authors string `json:"authors"`
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
func (m *ChangeBookTitleDelta) Key() string {
	return "название.изменено"
}
func (m *ChangeBookAuthorsDelta) Key() string {
	return "авторство.изменено"
}
