package types

import (
	"testing"

	"github.com/bzawada1/blockchain-and-go/crypto"
	pr "github.com/bzawada1/blockchain-and-go/proto"
	"github.com/bzawada1/blockchain-and-go/util"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	fromPrivKey := crypto.GeneratePrivateKey()
	fromAddr := fromPrivKey.Public().Address().Bytes()

	toPrivKey := crypto.GeneratePrivateKey()
	toAddr := toPrivKey.Public().Address().Bytes()

	input := &pr.TxInput{
		PrevTxHash:   util.RandomHash(),
		PrevOutIndex: 0,
		PublicKey:    fromPrivKey.Public().Bytes(),
	}
	output1 := &pr.TxOutput{
		Amount:  5,
		Address: toAddr,
	}
	output2 := &pr.TxOutput{
		Amount:  95,
		Address: fromAddr,
	}
	tx := &pr.Transaction{
		Version: 1,
		Inputs:  []*pr.TxInput{input},
		Outputs: []*pr.TxOutput{output1, output2},
	}

	sig := SignTransaction(fromPrivKey, tx)
	input.Signature = sig.Bytes()
	assert.True(t, VerifyTransaction(tx))
}
