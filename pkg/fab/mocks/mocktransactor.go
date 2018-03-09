/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mocks

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/common/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/context/api/fab"
	pb "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
)

// MockTransactor provides an implementation of Transactor that exposes all its context.
type MockTransactor struct {
	Ctx       context.Client
	ChannelID string
	Orderers  []fab.Orderer
}

// CreateTransactionHeader creates a Transaction Header based on the current context.
func (t *MockTransactor) CreateTransactionHeader() (fab.TransactionHeader, error) {
	return &MockTransactionHeader{}, nil
}

// SendTransactionProposal sends a TransactionProposal to the target peers.
func (t *MockTransactor) SendTransactionProposal(proposal *fab.TransactionProposal, targets []fab.ProposalProcessor) ([]*fab.TransactionProposalResponse, error) {
	response := make([]*fab.TransactionProposalResponse, 1, 1)
	response[0] = &fab.TransactionProposalResponse{Endorser: "example.com", Status: 99,
		ProposalResponse: &pb.ProposalResponse{Response: &pb.Response{Payload: []byte("abc")}},
	}
	return response, nil
}

// CreateTransaction create a transaction with proposal response.
func (t *MockTransactor) CreateTransaction(request fab.TransactionRequest) (*fab.Transaction, error) {
	response := &fab.Transaction{
		Proposal: &fab.TransactionProposal{
			Proposal: &pb.Proposal{},
		},
		Transaction: &pb.Transaction{},
	}
	return response, nil
}

// SendTransaction send a transaction to the chain’s orderer service (one or more orderer endpoints) for consensus and committing to the ledger.
func (t *MockTransactor) SendTransaction(tx *fab.Transaction) (*fab.TransactionResponse, error) {
	response := &fab.TransactionResponse{
		Orderer: "example.com",
	}
	return response, nil
}