package client

import (
	"context"
	"github.com/seracoin/serad/cmd/serawallet/daemon/server"
	"time"

	"github.com/pkg/errors"

	"github.com/seracoin/serad/cmd/serawallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the serawalletd server, and returns the client instance
func Connect(address string) (pb.SerawalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("serawallet daemon is not running, start it with `serawallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewSerawalletdClient(conn), func() {
		conn.Close()
	}, nil
}
