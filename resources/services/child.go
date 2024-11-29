package services

import (
	"context"
	"time"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func ChildTable() *schema.Table {
	return &schema.Table{
		Name:     "concurrency-childtable-child_table",
		Resolver: fetchChildTable,
		Columns: []schema.Column{
			{
				Name: "column",
				Type: arrow.BinaryTypes.String,
			},
		},
	}
}

func fetchChildTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	// cl := meta.(*client.Client)

	time.Sleep(500 * time.Millisecond)

	return nil
}
