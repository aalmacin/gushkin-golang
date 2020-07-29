package hooks

import (
	"fmt"

	"github.com/go-pg/pg"
)

type DBLogger struct{}

func (d DBLogger) BeforeQuery(q *pg.QueryEvent) {
}

func (d DBLogger) AfterQuery(p *pg.QueryEvent) {
	q, _ := p.FormattedQuery()
	fmt.Println(q)
}
