package contractapi

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// TransactionContextInterface defines the interface which TransactionContext
// meets. This can be taken by transacton functions on a contract which has not set
// a custom transaction context to allow transaction functions to take an interface
// to simplify unit testing.
type TransactionContextInterface interface {
	// GetStub should provide a way to access the stub set by Init/Invoke
	GetStub() shim.ChaincodeStubInterface
	Response(payload interface{}) peer.Response
	BadRequest(msg string) peer.Response
}

// SettableTransactionContextInterface defines functions a valid transaction context
// should have. Transaction context's set for contracts to be used in chaincode
// must implement this interface.
type SettableTransactionContextInterface interface {
	// SetStub should provide a way to pass the stub from a chaincode transaction
	// call to the transaction context so that it can be used by contract functions.
	// This is called by Init/Invoke with the stub passed.
	SetStub(shim.ChaincodeStubInterface)
}

// TransactionContext is a basic transaction context to be used in contracts,
// containing minimal required functionality use in contracts as part of
// chaincode. Provides access to the stub and clientIdentity of a transaction.
// If a contract implements the ContractInterface using the Contract struct then
// this is the default transaction context that will be used.
type TransactionContext struct {
	stub shim.ChaincodeStubInterface
}

// SetStub stores the passed stub in the transaction context
func (ctx *TransactionContext) SetStub(stub shim.ChaincodeStubInterface) {
	ctx.stub = stub
}

// GetStub returns the current set stub
func (ctx *TransactionContext) GetStub() shim.ChaincodeStubInterface {
	return ctx.stub
}

func (ctx *TransactionContext) Response(payload interface{}) peer.Response {
	return shim.Success(payload.([]byte))
}

func (ctx *TransactionContext) BadRequest(msg string) peer.Response {
	fmt.Println("400 response", msg)
	return peer.Response{
		Status:  400,
		Message: msg,
	}
}
