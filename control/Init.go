package control

import (
	"confirmchaincode/log"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type ProductTrace struct {
}

func (t *ProductTrace) Init(stub shim.ChaincodeStubInterface) peer.Response {
	log.Logger.Info("Init")
	return shim.Success(nil)
}
