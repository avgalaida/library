package domain

type BasedAggregate struct {
	ID        string `json:"id"`
	Meta      int    `json:"meta"`
	CreatedAt string `json:"createdAt"`
}
