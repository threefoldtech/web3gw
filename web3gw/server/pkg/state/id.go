package state

import (
	"golang.org/x/net/context"
)

const conUUID = "conUUID"

type labelContextKey struct{}

func IDFromContext(ctx context.Context) string {
	return ctx.Value(conUUID).(string)
}
