package node

import (
	"context"
	"fmt"

	pr "github.com/bzawada1/blockchain-and-go/proto"
	"google.golang.org/grpc/peer"
)

type Node struct {
	version string
	pr.UnimplementedNodeServer
}

func NewNode() *Node {
	return &Node{
		version: "blockchain-0.1",
	}
}

func (n *Node) HandleTransaction(ctx context.Context, tx *pr.Transaction) (*pr.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println(peer)
	return &pr.Ack{}, nil
}

func (n *Node) Handshake(ctx context.Context, version *pr.Version) (*pr.Version, error) {
	return &pr.Version{
		Version: n.version,
		Height:  100,
	}, nil
}
