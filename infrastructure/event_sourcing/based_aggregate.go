package event_sourcing

type BasedAggregate struct {
	ID        string `json:"id"`
	Meta      string `json:"meta"`
	CreatedAt string `json:"createdAt"`
}
