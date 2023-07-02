package core

import (
	"context"

	"github.com/hazelcast/hazelcast-go-client"
)

func T1(ctx context.Context) {
	hazelcast.StartNewClientWithConfig(ctx, hazelcast.NewConfig())
}
