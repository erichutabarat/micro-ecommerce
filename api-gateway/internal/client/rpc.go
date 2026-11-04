package client

import (
	"context"
	"net"
	"net/rpc"
	"time"
)

const dialTimeout = 3 * time.Second

func Call(ctx context.Context, address, method string, args, reply any) error {
	conn, err := net.DialTimeout("tcp", address, dialTimeout)
	if err != nil {
		return err
	}

	c := rpc.NewClient(conn)
	defer c.Close()

	call := c.Go(method, args, reply, make(chan *rpc.Call, 1))

	select {
	case <-ctx.Done():
		return ctx.Err()
	case res := <-call.Done:
		return res.Error
	}
}
