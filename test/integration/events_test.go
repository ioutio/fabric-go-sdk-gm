/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package integration

import (
	"testing"
	"time"

	api "github.com/hyperledger/fabric-sdk-go/api"

	"github.com/hyperledger/fabric-sdk-go/pkg/util"
	"github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"
)

const (
	eventTimeout = time.Second * 30
)

func TestEvents(t *testing.T) {
	testSetup := initializeTests(t)

	testFailedTx(t, testSetup)
	testFailedTxErrorCode(t, testSetup)
	testReconnectEventHub(t, testSetup)
	testMultipleBlockEventCallbacks(t, testSetup)
}

func initializeTests(t *testing.T) BaseSetupImpl {
	testSetup := BaseSetupImpl{
		ConfigFile:      "../fixtures/config/config_test.yaml",
		ChannelID:       "mychannel",
		ChannelConfig:   "../fixtures/channel/mychannel.tx",
		ConnectEventHub: true,
	}

	if err := testSetup.Initialize(); err != nil {
		t.Fatalf(err.Error())
	}

	testSetup.ChainCodeID = util.GenerateRandomID()

	// Install and Instantiate Events CC
	if err := testSetup.InstallCC(testSetup.ChainCodeID, "github.com/events_cc", "v0", nil); err != nil {
		t.Fatalf("installCC return error: %v", err)
	}

	if err := testSetup.InstantiateCC(testSetup.ChainCodeID, testSetup.ChannelID, "github.com/events_cc", "v0", nil); err != nil {
		t.Fatalf("instantiateCC return error: %v", err)
	}

	return testSetup
}

func testFailedTx(t *testing.T, testSetup BaseSetupImpl) {
	// Arguments for events CC
	var args []string
	args = append(args, "invoke")
	args = append(args, "invoke")
	args = append(args, "SEVERE")

	tpResponses1, tx1, err := util.CreateAndSendTransactionProposal(testSetup.Channel, testSetup.ChainCodeID, testSetup.ChannelID, args, []api.Peer{testSetup.Channel.GetPrimaryPeer()}, nil)
	if err != nil {
		t.Fatalf("CreateAndSendTransactionProposal return error: %v \n", err)
	}

	tpResponses2, tx2, err := util.CreateAndSendTransactionProposal(testSetup.Channel, testSetup.ChainCodeID, testSetup.ChannelID, args, []api.Peer{testSetup.Channel.GetPrimaryPeer()}, nil)
	if err != nil {
		t.Fatalf("CreateAndSendTransactionProposal return error: %v \n", err)
	}

	// Register tx1 and tx2 for commit/block event(s)
	done1, fail1 := util.RegisterTxEvent(tx1, testSetup.EventHub)
	defer testSetup.EventHub.UnregisterTxEvent(tx1)

	done2, fail2 := util.RegisterTxEvent(tx2, testSetup.EventHub)
	defer testSetup.EventHub.UnregisterTxEvent(tx2)

	go monitorFailedTx(t, testSetup, done1, fail1, done2, fail2)

	// Test invalid transaction: create 2 invoke requests in quick succession that modify
	// the same state variable which should cause one invoke to be invalid
	_, err = util.CreateAndSendTransaction(testSetup.Channel, tpResponses1)
	if err != nil {
		t.Fatalf("First invoke failed err: %v", err)
	}
	_, err = util.CreateAndSendTransaction(testSetup.Channel, tpResponses2)
	if err != nil {
		t.Fatalf("Second invoke failed err: %v", err)
	}

}

func monitorFailedTx(t *testing.T, testSetup BaseSetupImpl, done1 chan bool, fail1 chan error, done2 chan bool, fail2 chan error) {
	rcvDone := false
	rcvFail := false
	timeout := time.After(eventTimeout)

Loop:
	for !rcvDone || !rcvFail {
		select {
		case <-done1:
			rcvDone = true
		case <-fail1:
			t.Fatalf("Received fail for first invoke")
		case <-done2:
			t.Fatalf("Received success for second invoke")
		case <-fail2:
			rcvFail = true
		case <-timeout:
			t.Logf("Timeout: Didn't receive events")
			break Loop
		}
	}

	if !rcvDone || !rcvFail {
		t.Fatalf("Didn't receive events (done: %t; fail %t)", rcvDone, rcvFail)
	}
}

func testFailedTxErrorCode(t *testing.T, testSetup BaseSetupImpl) {
	// Arguments for events CC
	var args []string
	args = append(args, "invoke")
	args = append(args, "invoke")
	args = append(args, "SEVERE")

	tpResponses1, tx1, err := util.CreateAndSendTransactionProposal(testSetup.Channel, testSetup.ChainCodeID, testSetup.ChannelID, args, []api.Peer{testSetup.Channel.GetPrimaryPeer()}, nil)
	if err != nil {
		t.Fatalf("CreateAndSendTransactionProposal return error: %v \n", err)
	}

	tpResponses2, tx2, err := util.CreateAndSendTransactionProposal(testSetup.Channel, testSetup.ChainCodeID, testSetup.ChannelID, args, []api.Peer{testSetup.Channel.GetPrimaryPeer()}, nil)
	if err != nil {
		t.Fatalf("CreateAndSendTransactionProposal return error: %v \n", err)
	}

	done := make(chan bool)
	fail := make(chan pb.TxValidationCode)

	testSetup.EventHub.RegisterTxEvent(tx1, func(txId string, errorCode pb.TxValidationCode, err error) {
		if err != nil {
			fail <- errorCode
		} else {
			done <- true
		}
	})

	defer testSetup.EventHub.UnregisterTxEvent(tx1)

	done2 := make(chan bool)
	fail2 := make(chan pb.TxValidationCode)

	testSetup.EventHub.RegisterTxEvent(tx2, func(txId string, errorCode pb.TxValidationCode, err error) {
		if err != nil {
			fail2 <- errorCode
		} else {
			done2 <- true
		}
	})

	defer testSetup.EventHub.UnregisterTxEvent(tx2)

	go monitorFailedTxErrorCode(t, testSetup, done, fail, done2, fail2)

	// Test invalid transaction: create 2 invoke requests in quick succession that modify
	// the same state variable which should cause one invoke to be invalid
	_, err = util.CreateAndSendTransaction(testSetup.Channel, tpResponses1)
	if err != nil {
		t.Fatalf("First invoke failed err: %v", err)
	}
	_, err = util.CreateAndSendTransaction(testSetup.Channel, tpResponses2)
	if err != nil {
		t.Fatalf("Second invoke failed err: %v", err)
	}
}

func monitorFailedTxErrorCode(t *testing.T, testSetup BaseSetupImpl, done chan bool, fail chan pb.TxValidationCode, done2 chan bool, fail2 chan pb.TxValidationCode) {
	rcvDone := false
	rcvFail := false
	timeout := time.After(eventTimeout)

Loop:
	for !rcvDone || !rcvFail {
		select {
		case <-done:
			rcvDone = true
		case <-fail:
			t.Fatalf("Received fail for first invoke")
		case <-done2:
			t.Fatalf("Received success for second invoke")
		case errorValidationCode := <-fail2:
			if errorValidationCode.String() != "MVCC_READ_CONFLICT" {
				t.Fatalf("Expected error code MVCC_READ_CONFLICT. Got %s", errorValidationCode.String())
			}
			rcvFail = true
		case <-timeout:
			t.Logf("Timeout: Didn't receive events")
			break Loop
		}
	}

	if !rcvDone || !rcvFail {
		t.Fatalf("Didn't receive events (done: %t; fail %t)", rcvDone, rcvFail)
	}
}

func testReconnectEventHub(t *testing.T, testSetup BaseSetupImpl) {
	// Test disconnect event hub
	testSetup.EventHub.Disconnect()
	if testSetup.EventHub.IsConnected() {
		t.Fatalf("Failed to disconnect event hub")
	}

	// Reconnect event hub
	if err := testSetup.EventHub.Connect(); err != nil {
		t.Fatalf("Failed to connect event hub")
	}
}

func testMultipleBlockEventCallbacks(t *testing.T, testSetup BaseSetupImpl) {
	// Arguments for events CC
	var args []string
	args = append(args, "invoke")
	args = append(args, "invoke")
	args = append(args, "SEVERE")

	// Create and register test callback that will be invoked upon block event
	test := make(chan bool)
	testSetup.EventHub.RegisterBlockEvent(func(block *common.Block) {
		t.Logf("Received test callback on block event")
		test <- true
	})

	tpResponses, tx, err := util.CreateAndSendTransactionProposal(testSetup.Channel, testSetup.ChainCodeID, testSetup.ChannelID, args, []api.Peer{testSetup.Channel.GetPrimaryPeer()}, nil)
	if err != nil {
		t.Fatalf("CreateAndSendTransactionProposal returned error: %v \n", err)
	}

	// Register tx for commit/block event(s)
	done, fail := util.RegisterTxEvent(tx, testSetup.EventHub)
	defer testSetup.EventHub.UnregisterTxEvent(tx)

	go monitorMultipleBlockEventCallbacks(t, testSetup, done, fail, test)

	_, err = util.CreateAndSendTransaction(testSetup.Channel, tpResponses)
	if err != nil {
		t.Fatalf("CreateAndSendTransaction failed with error: %v", err)
	}
}

func monitorMultipleBlockEventCallbacks(t *testing.T, testSetup BaseSetupImpl, done chan bool, fail chan error, test chan bool) {
	rcvBlock := false
	rcvTx := false
	timeout := time.After(eventTimeout)

Loop:
	for !rcvTx || !rcvBlock {
		select {
		case <-done:
			rcvTx = true
		case <-fail:
			t.Fatalf("Received tx failure")
		case <-test:
			rcvBlock = true
		case <-timeout:
			t.Logf("Timeout while waiting for events")
			break Loop
		}
	}

	if !rcvTx || !rcvBlock {
		t.Fatalf("Didn't receive events (tx: %t; block %t)", rcvTx, rcvBlock)
	}
}
