package util

import (
	randc "crypto/rand"
	"github.com/bzawada1/blockchain-and-go/proto"
	"io"
	"math/rand"
	"time"
)

func RandomHash() []byte {
	hash := make([]byte, 32)
	io.ReadFull(randc.Reader, hash)
	return hash
}

func RandomBlock() *proto.Block {
	header := &proto.Header{
		Version:   1,
		Height:    int32(rand.Intn(1000)),
		PrevHash:  RandomHash(),
		RootHash:  RandomHash(),
		Timestamp: time.Now().Unix(),
	}
	transactions := &proto.Transaction{}

	return &proto.Block{
		Header:       header,
		Transactions: []*proto.Transaction{transactions},
	}
}
