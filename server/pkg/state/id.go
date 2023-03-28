package state

import "golang.org/x/net/context"

const jrpcID = "jrpc-uuid"

func IDFromContext(ctx context.Context) string {
	return ctx.Value(jrpcID).(string)
}
