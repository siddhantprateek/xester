// Package chaincode contains the code necessary for chaincode to interact
// with a Hyperledger Fabric peer.
package chaincode

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

// Fabric enforces that every chaincode must implement Chaincode interface.
type REState struct {
	contractapi.Contract
	Name               string
	Info               metadata.InfoMetadata
	UnknownTransaction interface{}
	BeforeTransaction  interface{}
	AfterTransaction   interface{}
}

type Property struct {
	propertyLocation string
	price            int
	sellerAddress    string
	buyerAddress     string
	isSold           bool
	isForSale        bool
}

// Init: called when the chaincode is instantiated by the blockchain network
func (c *REState) Init(stub shim.ChaincodeStubInterface) pb.Response {

	return shim.Success([]byte("Asset modified"))
}

// Invoke: called when the client invokes a specific function to process
// the transaction proposal
func (c *REState) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success([]byte("Invoke"))
}

// To create the transaction context for the function defined within the smart contract
// that read and write data to the blockchain ledger.
func (c *REState) InitLedger(ctx contractapi.TransactionContextInterface) error {
	properties := []Property{
		{
			propertyLocation: "Alabama",
			price:            450000,
			sellerAddress:    "Alabama",
			buyerAddress:     "Texas",
			isSold:           false,
			isForSale:        true,
		},
		{
			propertyLocation: "New York",
			price:            2100000,
			sellerAddress:    "Kentucky",
			buyerAddress:     "New Jersey",
			isSold:           false,
			isForSale:        true,
		},
	}

	for id, prop_ty := range properties {
		propertyAsBytes, _ := json.Marshal(prop_ty)
		err := ctx.GetStub().PutState("PROPERTY"+strconv.Itoa(id), propertyAsBytes)
		if err != nil {
			return fmt.Errorf("Failed to put to world state %s", err.Error())
		}
	}
	return nil
}

func (c *REState) CreateProperty(
	ctx contractapi.TransactionContextInterface,
	propId string,
	propLocation string,
	propPrice int,
	sellerAddr string,
	buyerAddr string,
	soldStatus bool,
	isForSaleStatus bool) error {
	property := Property{
		propertyLocation: propLocation,
		price:            propPrice,
		sellerAddress:    sellerAddr,
		buyerAddress:     buyerAddr,
		isSold:           soldStatus,
		isForSale:        isForSaleStatus,
	}

	propAsBytes, _ := json.Marshal(property)
	return ctx.GetStub().PutState(propId, propAsBytes)
}
