// Package chaincode contains the code necessary for chaincode to interact
// with a Hyperledger Fabric peer.
package chaincode

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

// Fabric enforces that every chaincode must implement Chaincode interface.
type AppChaincode struct {
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
