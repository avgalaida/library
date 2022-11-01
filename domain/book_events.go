package domain

type IDelta interface {
	Key() string
	ApplyOn(b *Book)
}

type CreateBookDelta struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	Title     string `json:"title"`
	Authors   string `json:"authors"`
	CreatedAt string `json:"createdAt"`
	Apply     func()
}

func (d *CreateBookDelta) ApplyOn(b *Book) {
	b.Status = d.Status
	b.Title = d.Title
	b.Authors = d.Authors
}

type DeleteBookDelta struct {
	ID string `json:"id"`
}

func (d *DeleteBookDelta) ApplyOn(b *Book) {
	b.Status = "Недоступна"
}

type RestoreBookDelta struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (d *RestoreBookDelta) ApplyOn(b *Book) {
	b.Status = d.Status
}

type ChangeBookTitleDelta struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (d *ChangeBookTitleDelta) ApplyOn(b *Book) {
	b.Title = d.Title
}

type ChangeBookAuthorsDelta struct {
	ID      string `json:"id"`
	Authors string `json:"authors"`
}

func (d *ChangeBookAuthorsDelta) ApplyOn(b *Book) {
	b.Authors = d.Authors
}

type RollbackBookDelta struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Title   string `json:"title"`
	Authors string `json:"authors"`
}

func (d *RollbackBookDelta) ApplyOn(b *Book) {
	b.Title = d.Title
	b.Authors = d.Authors
	b.Status = d.Status
}

func (d *CreateBookDelta) Key() string {
	return "книга.создана"
}
func (d *DeleteBookDelta) Key() string {
	return "книга.удалена"
}
func (d *RestoreBookDelta) Key() string {
	return "книга.восстановлена"
}
func (d *ChangeBookTitleDelta) Key() string {
	return "название.изменено"
}
func (d *ChangeBookAuthorsDelta) Key() string {
	return "авторство.изменено"
}
func (d *RollbackBookDelta) Key() string {
	return "откат.версии"
}
