package hooks

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
)

type DebugHook struct {
	// Verbose causes hook to print all queries (even those without an error).
	Verbose   bool
	EmptyLine bool
	l         *zap.Logger
}

func NewDebugHook(l *zap.Logger) *DebugHook {
	return &DebugHook{
		l: l,
	}
}

var _ pg.QueryHook = (*DebugHook)(nil)

func (h *DebugHook) BeforeQuery(ctx context.Context, evt *pg.QueryEvent) (context.Context, error) {
	q, err := evt.FormattedQuery()
	if err != nil {
		return nil, err
	}

	h.l.Named("[PG-INSTANCE]").Debug("executing a query...",
		zap.Time("at", evt.StartTime),
		zap.ByteString("query", q),
	)

	if evt.Err != nil {
		fmt.Printf("%s executing a query:\n%s\n", evt.Err, q)
	} else if h.Verbose {
		if h.EmptyLine {
			fmt.Println()
		}
		fmt.Println(string(q))
	}

	return ctx, nil
}

func (h *DebugHook) AfterQuery(c context.Context, evt *pg.QueryEvent) error {
	return nil
}
