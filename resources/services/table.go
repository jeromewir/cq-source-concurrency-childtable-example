package services

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

type item struct {
	ID string
}

func SampleTable() *schema.Table {
	return &schema.Table{
		Name:     "concurrency-childtable-example_sample_table",
		Resolver: fetchSampleTable,
		Columns: []schema.Column{
			{
				Name: "id",
				Type: arrow.BinaryTypes.String,
			},
		},
		Relations: []*schema.Table{
			ChildTable(),
		},
	}
}

func fetchSampleTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	// cl := meta.(*client.Client)

	items := []item{}

	// Adding all of them at once
	for i := range 10 {
		items = append(items, item{ID: fmt.Sprintf("id-%d", i)})
	}

	res <- items

	// Adding them one by one
	for _, item := range items {
		res <- item
	}

	return nil
}
