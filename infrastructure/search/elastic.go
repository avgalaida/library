package search

import (
	"bytes"
	"encoding/json"
	"github.com/avgalaida/library/domain"
	elastic "github.com/elastic/go-elasticsearch/v7"
)

type ElasticRepository struct {
	client *elastic.Client
}

func NewElastic(url string) (*ElasticRepository, error) {
	client, err := elastic.NewClient(elastic.Config{
		Addresses: []string{url},
	})
	if err != nil {
		return nil, err
	}
	_, err = client.Info()
	if err != nil {
		return nil, err
	}
	return &ElasticRepository{client}, nil
}

func (r *ElasticRepository) Close() {
}

func (r *ElasticRepository) InsertBook(book domain.BookView) {
	body, _ := json.Marshal(book)
	r.client.Index(
		"books",
		bytes.NewReader(body),
		r.client.Index.WithDocumentID(book.ID),
		r.client.Index.WithRefresh("wait_for"),
	)
}

func (r *ElasticRepository) SearchBooks(query string, skip, take uint64) (result []domain.BookView) {
	var buf bytes.Buffer
	reqBody := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":            query,
				"fields":           []string{"body"},
				"fuzziness":        3,
				"cutoff_frequency": 0.0001,
			},
		},
	}
	json.NewEncoder(&buf).Encode(reqBody)

	res, _ := r.client.Search(
		r.client.Search.WithIndex("books"),
		r.client.Search.WithFrom(int(skip)),
		r.client.Search.WithSize(int(take)),
		r.client.Search.WithBody(&buf),
		r.client.Search.WithTrackTotalHits(true),
	)

	defer res.Body.Close()

	type Response struct {
		Took int64
		Hits struct {
			Total struct {
				Value int64
			}
			Hits []*struct {
				Source domain.BookView `json:"_source"`
			}
		}
	}
	resBody := Response{}
	json.NewDecoder(res.Body).Decode(&resBody)
	var books []domain.BookView
	for _, hit := range resBody.Hits.Hits {
		books = append(books, hit.Source)
	}
	return books
}
