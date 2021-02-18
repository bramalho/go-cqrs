package search

import (
	"context"
	"encoding/json"
	"github.com/bramalho/go-cqrs/model"
	"github.com/olivere/elastic"
	"log"
)

type ElasticRepository struct {
	client *elastic.Client
}

func NewElastic(url string) (*ElasticRepository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	return &ElasticRepository{client}, nil
}

func (r *ElasticRepository) Close() {
}

func (r *ElasticRepository) InsertTodo(ctx context.Context, todo model.Todo) error {
	_, err := r.client.Index().
		Index("todos").
		Type("todo").
		Id(todo.ID).
		BodyJson(todo).
		Refresh("wait_for").
		Do(ctx)
	return err
}

func (r *ElasticRepository) SearchTodos(ctx context.Context, query string, skip uint64, take uint64) ([]model.Todo, error) {
	result, err := r.client.Search().
		Index("todos").
		Query(
			elastic.NewMultiMatchQuery(query, "body").
				Fuzziness("3").
				PrefixLength(1).
				CutoffFrequency(0.0001),
		).
		From(int(skip)).
		Size(int(take)).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	todos := []model.Todo{}
	for _, hit := range result.Hits.Hits {
		log.Println(hit.Source)
		var todo model.Todo
		if err = json.Unmarshal(*hit.Source, &todo); err != nil {
			log.Println(err)
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
