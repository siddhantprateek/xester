package main

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	chaincode "github.com/siddhantprateek/xester/chaincode"
)

func main() {

	// logger.SetLevel(shim.LogInfo)
	err := shim.Start(new(chaincode.AppChaincode))
	if err != nil {
		fmt.Println("Error starting Chaincode", err.Error())
	}
}
