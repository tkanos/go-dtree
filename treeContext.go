package dtree

import (
	"context"
	"fmt"
)

// contextKey Log Context Key
type contextKey string

const descriptionContext = contextKey("nodes-description")

// contextValue add value context on ctx
func contextValue(ctx context.Context, id int, key string, value interface{}, operator string, tvalue interface{}) context.Context {
	description := fmt.Sprintf("%d : %s %v %s %v", id, key, value, operator, tvalue)
	if v := ctx.Value(descriptionContext); v != nil {
		s, ok := v.([]string)
		if ok {
			s = append(s, description)
			ctx = context.WithValue(ctx, descriptionContext, s)
		}
	} else {
		ctx = context.WithValue(ctx, descriptionContext, []string{description})
	}

	return ctx
}

// GetNodePathFromContext gets the node path from the context
func GetNodePathFromContext(ctx context.Context) []string {
	if v := ctx.Value(descriptionContext); v != nil {
		s, ok := v.([]string)
		if ok {
			return s
		}
	}

	return nil
}
