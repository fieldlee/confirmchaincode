package main

import (
	"confirmchaincode/control"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(control.ProductTrace))
	if err != nil {
		fmt.Printf("Error starting ProductTrace: %s", err)
	}
}
