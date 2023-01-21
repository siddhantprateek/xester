// Package chaincode contains the code necessary for chaincode to interact
// with a Hyperledger Fabric peer.
package chaincode

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

// Fabric enforces that every chaincode must implement Chaincode interface.
type AppChaincode struct {
	contractapi.Contract
	Name               string
	Info               metadata.InfoMetadata
	UnknownTransaction interface{}
	BeforeTransaction  interface{}
	AfterTransaction   interface{}
}

// Init: called when the chaincode is instantiated by the blockchain network
func (c *AppChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success([]byte("Asset modified"))
}

// Invoke: called when the client invokes a specific function to process
// the transaction proposal
func (c *AppChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success([]byte("Invoke"))
}

// To create the transaction context for the function defined within the smart contract
// that read and write data to the blockchain ledger.
func (c *AppChaincode) Create(ctx contractapi.TransactionContextInterface) error {

	return nil
}
